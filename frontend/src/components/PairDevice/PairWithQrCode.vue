<script setup lang="ts">
import ErrorBanner from "@/components/ErrorBanner.vue";
import { preferredServiceHost } from "@/helpers/mdns";
import { useAppStore } from "@/store";
import { nanoid } from "nanoid";
import QrcodeVue from "qrcode.vue";
import { computed, onBeforeMount, ref, watch } from "vue";

const emit = defineEmits<{
  pair: [];
}>();

const name = ref("");
const password = ref("");
const pairing = ref(false);

const error = ref("");

onBeforeMount(() => {
  name.value = "ADB_WIFI_" + nanoid();
  password.value = nanoid();
});

const store = useAppStore();

watch(store.services, async (services) => {
  const pairingService = [...services.values()].find(
    (svc) => svc.name === name.value,
  );

  if (!pairingService) {
    return;
  }

  error.value = "";
  pairing.value = true;

  try {
    await store.pairDevice(
      preferredServiceHost(pairingService),
      password.value,
    );
    emit("pair");
  } catch (e) {
    error.value = e as string;
  } finally {
    pairing.value = false;
  }
});

const qrContent = computed(
  () => `WIFI:T:ADB;S:${name.value};P:${password.value};;`,
);
</script>

<template>
  <div class="column no-wrap items-center">
    <div class="col column">
      <ol>
        <li>
          Go to
          <strong>Developer options</strong>
          &rarr;
          <strong>Wireless debugging</strong>
        </li>
        <li>
          Select
          <strong> Pair device with QR code </strong>
        </li>
        <li>Scan this QR-code:</li>
      </ol>
      <div class="self-center relative-position">
        <qrcode-vue :value="qrContent" :size="256" />
        <q-inner-loading :showing="pairing" />
      </div>
    </div>
    <error-banner
      label="Pairing error:"
      v-model:error="error"
      class="self-stretch"
    />
  </div>
</template>
