<script lang="ts" setup>
import AppSettings from "@/components/AppSettings.vue";
import DevicesList from "@/components/DevicesList.vue";
import PairDevice from "@/components/PairDevice.vue";
import { useAppStore } from "@/store";

useAppStore();

function styleFn(offset: number, height: number) {
  // Based on QPage.style: uses `height` instead of `minHeight`
  return {
    height:
      height === 0
        ? offset === 0
          ? "100vh"
          : `calc(100vh - ${offset}px)`
        : height - offset + "px",
  };
}
</script>

<template>
  <q-layout view="hHh lpR fFf">
    <q-header>
      <q-toolbar>
        <pair-device
          flat
          icon="mdi-handshake"
          no-caps
          dense
          label="Pair device"
        />

        <q-space />

        <app-settings flat round icon="mdi-cog" title="Settings" />
      </q-toolbar>
    </q-header>

    <q-page-container>
      <q-page :style-fn="styleFn">
        <devices-list class="full-height" />
      </q-page>
    </q-page-container>
  </q-layout>
</template>
