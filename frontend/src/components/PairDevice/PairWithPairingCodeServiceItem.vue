<script setup lang="ts">
import ErrorBanner from "@/components/ErrorBanner.vue";
import type { zeroconf } from "@/go/models";
import { preferredServiceHost } from "@/helpers/mdns";
import { useAppStore } from "@/store";
import { isIPv4 } from "is-ip";
import { computed, shallowRef } from "vue";

const { service } = defineProps<{
  service: zeroconf.Service;
}>();

const emit = defineEmits<{
  pair: [];
}>();

const addr = computed(() => (service.addrs as string[]).find(isIPv4));

const pairingCode = shallowRef("");
const isPairing = shallowRef(false);

const error = shallowRef("");

const store = useAppStore();

async function doPair() {
  if (isPairing.value || !codeIsValid.value) {
    return;
  }
  isPairing.value = true;
  error.value = "";

  try {
    await store.pairDevice(
      `${preferredServiceHost(service)}:${service.port}`,
      pairingCode.value,
    );
    emit("pair");
  } catch (e) {
    error.value = e as string;
  } finally {
    isPairing.value = false;
  }
}

const codeIsValid = computed(() => /^[0-9]{6}$/.test(pairingCode.value));
</script>

<template>
  <q-item>
    <q-item-section>
      <q-item-label> {{ addr }}:{{ service.port }} </q-item-label>
      <q-item-label caption class="ellipsis-2-lines" :title="service.name">
        {{ service.name }}
      </q-item-label>
    </q-item-section>
    <q-item-section>
      <q-input
        v-model="pairingCode"
        dense
        label="6-digit pairing code"
        mask="######"
        :disable="isPairing"
        @keydown.enter="doPair"
      >
        <template #append>
          <q-btn
            round
            dense
            flat
            icon="mdi-link"
            color="primary"
            :loading="isPairing"
            :error="!!error"
            @click="doPair"
          />
        </template>
      </q-input>
    </q-item-section>
  </q-item>
  <error-banner label="Pairing error:" v-model:error="error" />
</template>

<style scoped></style>
