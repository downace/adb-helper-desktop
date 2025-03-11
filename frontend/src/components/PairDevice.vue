<script setup lang="ts">
import { shallowRef } from "vue";
import PairWithPairingCode from "./PairDevice/PairWithPairingCode.vue";
import PairWithQrCode from "./PairDevice/PairWithQrCode.vue";

const isOpen = shallowRef(false);

const tab = shallowRef("qr");
</script>

<template>
  <q-btn v-bind="$attrs" @click="isOpen = true" />
  <q-dialog
    v-model="isOpen"
    maximized
    transition-show="slide-up"
    transition-hide="slide-down"
  >
    <q-card class="column">
      <q-toolbar class="bg-primary text-white">
        <q-toolbar-title>Pair Device</q-toolbar-title>
        <q-btn flat round icon="mdi-close" @click="isOpen = false" />
      </q-toolbar>
      <q-tabs v-model="tab" no-caps inline-label>
        <q-tab name="qr" icon="mdi-qrcode-scan" label="QR Code" />
        <q-tab
          name="code"
          icon="mdi-form-textbox-password"
          label="Pairing Code"
        />
      </q-tabs>
      <q-tab-panels v-model="tab" class="col full-width">
        <q-tab-panel name="qr" class="no-padding">
          <pair-with-qr-code class="full-height" @pair="isOpen = false" />
        </q-tab-panel>
        <q-tab-panel name="code" class="no-padding">
          <pair-with-pairing-code class="full-height" @pair="isOpen = false" />
        </q-tab-panel>
      </q-tab-panels>
    </q-card>
  </q-dialog>
</template>

<style scoped>
:deep(.q-tab-panel.no-padding) {
  padding: 0;
}
</style>
