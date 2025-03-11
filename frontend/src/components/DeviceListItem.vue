<script setup lang="ts">
import AdbDeviceStateIcon from "@/components/AdbDeviceStateIcon.vue";
import ErrorBanner from "@/components/ErrorBanner.vue";
import { StartShell } from "@/go/main/App";
import { preferredServiceHost } from "@/helpers/mdns";
import { useAppStore } from "@/store";
import { Device } from "@/types";
import { isIPv4 } from "is-ip";
import { computed, shallowRef } from "vue";

const store = useAppStore();

const { device } = defineProps<{
  device: Device;
}>();

const deviceName = computed(() => {
  if (device.service) {
    const addr =
      ((device.service.addrs ?? []) as string[]).find(isIPv4) ??
      device.service.hostname;
    return `${addr}:${device.service.port}`;
  } else if (device.adbDevice) {
    return device.adbDevice.id;
  }

  return "";
});

const connectError = shallowRef("");
const shellError = shallowRef("");

async function toggleConnection() {
  connectError.value = "";
  connecting.value = true;
  try {
    if (device.adbDevice) {
      await store.disconnectDevice(device.adbDevice.id);
    } else if (device.service) {
      await store.connectDevice(preferredServiceHost(device.service));
    }
  } catch (e) {
    connectError.value = e as string;
  } finally {
    connecting.value = false;
  }
}

const isConnected = computed(() => device.adbDevice?.state === "device");

const connecting = shallowRef(false);

const adbDeviceUserFriendlyName = computed(() => {
  if (
    !device.adbDevice ||
    (!device.adbDevice!.brand && !device.adbDevice!.model)
  ) {
    return null;
  }

  return `${device.adbDevice.brand} ${device.adbDevice.model}`.trim() || null;
});

async function startShell() {
  shellError.value = "";
  try {
    await StartShell(device.adbDevice!.id);
  } catch (e) {
    shellError.value = e as string;
  }
}
</script>

<template>
  <q-item>
    <q-item-section avatar>
      <adb-device-state-icon :device="device.adbDevice" />
    </q-item-section>
    <q-item-section>
      <q-item-label :class="{ 'text-italic': !deviceName }">
        {{ deviceName || "Unknown device" }}
      </q-item-label>
      <q-item-label
        v-if="adbDeviceUserFriendlyName || device.adbDevice?.description.model"
        caption
      >
        {{ adbDeviceUserFriendlyName ?? device.adbDevice?.description.model }}
        <q-btn
          v-if="!adbDeviceUserFriendlyName"
          size="xs"
          dense
          flat
          round
          icon="mdi-help-circle"
        >
          <q-tooltip>
            Device brand and model is unknown, go to Settings to update known
            devices list
          </q-tooltip>
        </q-btn>
      </q-item-label>
    </q-item-section>
    <q-item-section v-if="isConnected" side>
      <q-btn flat round icon="mdi-dots-vertical" title="Actions">
        <q-menu auto-close>
          <q-list>
            <q-item clickable @click="startShell">
              <q-item-section avatar>
                <q-icon name="mdi-console" />
              </q-item-section>
              <q-item-section>
                <q-item-label> Start shell </q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </q-menu>
      </q-btn>
    </q-item-section>
    <q-item-section side>
      <q-btn flat round icon="mdi-information-outline" title="More info">
        <q-menu>
          <q-list separator>
            <template v-if="device.service">
              <q-item>
                <q-item-section>
                  <q-item-label caption>Service Name</q-item-label>
                  <q-item-label>
                    {{ device.service.name }}
                  </q-item-label>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label caption>Service Hostname</q-item-label>
                  <q-item-label>
                    {{ device.service.hostname }}
                  </q-item-label>
                </q-item-section>
              </q-item>
              <q-item v-if="device.service.addrs">
                <q-item-section>
                  <q-item-label caption>Service IPs</q-item-label>
                  <q-item-label>
                    <q-chip v-for="addr in device.service.addrs" dense square>
                      {{ addr }}
                    </q-chip>
                  </q-item-label>
                </q-item-section>
              </q-item>
              <q-item>
                <q-item-section>
                  <q-item-label caption>Service Port</q-item-label>
                  <q-item-label>
                    {{ device.service.port }}
                  </q-item-label>
                </q-item-section>
              </q-item>
            </template>
            <q-item v-if="device.adbDevice">
              <q-item-section>
                <q-item-label caption>ADB device description</q-item-label>
                <q-item-label>
                  <q-chip
                    v-for="(value, key) in device.adbDevice.description"
                    dense
                    square
                  >
                    {{ key }}:{{ value }}
                  </q-chip>
                </q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </q-menu>
      </q-btn>
    </q-item-section>
    <q-item-section side>
      <q-btn
        @click="toggleConnection"
        :loading="connecting"
        flat
        round
        :icon="isConnected ? 'mdi-link-off' : 'mdi-link'"
        color="primary"
        :title="isConnected ? 'Disconnect' : 'Connect'"
      />
    </q-item-section>
  </q-item>
  <error-banner label="Connection error" v-model:error="connectError" />
  <q-banner v-if="connectError" class="bg-primary text-white">
    TIP: Ensure that device is paired
  </q-banner>
  <error-banner label="Shell error" v-model:error="shellError" />
  <q-banner v-if="shellError" class="bg-primary text-white">
    TIP: Try specifying terminal command in Settings
  </q-banner>
</template>
