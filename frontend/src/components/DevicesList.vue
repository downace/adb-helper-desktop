<script setup lang="ts">
import { useAppStore } from "@/store";
import { Device } from "@/types";
import { computed } from "vue";
import DeviceListItem from "./DeviceListItem.vue";

const store = useAppStore();

const devices = computed<Device[]>(() => {
  const discoveredDevices: Device[] = [...store.services.values()]
    .filter((s) => s.type.name === "_adb-tls-connect._tcp")
    .map((service) => ({
      service,
      adbDevice: store.getConnectedDevice(service),
    }));

  const connectedDevices: Device[] = [...store.connectedDevices.values()]
    .filter((ad) => discoveredDevices.every((d) => d.adbDevice?.id !== ad.id))
    .map((adbDevice) => ({ service: null, adbDevice }));
  return [...discoveredDevices, ...connectedDevices];
});
</script>

<template>
  <q-list v-if="devices.length > 0" separator>
    <device-list-item v-for="device in devices" :device="device" />
  </q-list>
  <div v-else class="column no-wrap justify-center items-center q-pa-sm">
    <h6>No devices</h6>
    <div class="text-grey">Tips:</div>
    <ul class="text-grey">
      <li>
        Ensure that Android device is connected to the same Wi-Fi network as
        this PC
      </li>
      <li>Ensure that <strong>Wireless Debugging</strong> is enabled</li>
      <li>Try re-enabling <strong>Wireless Debugging</strong></li>
    </ul>
  </div>
</template>
