<script setup lang="ts">
import AppSettingsAdb from "@/components/AppSettings/AppSettingsAdb.vue";
import AppSettingsKnownDevices from "@/components/AppSettings/AppSettingsKnownDevices.vue";
import AppSettingsShell from "@/components/AppSettings/AppSettingsShell.vue";
import { shallowRef } from "vue";

const isOpen = shallowRef(false);
</script>

<template>
  <q-btn v-bind="$attrs" @click="isOpen = true" />
  <q-dialog
    v-model="isOpen"
    maximized
    transition-show="slide-up"
    transition-hide="slide-down"
  >
    <q-card class="column no-wrap">
      <q-toolbar class="bg-primary text-white">
        <q-toolbar-title>Settings</q-toolbar-title>
        <q-btn flat round icon="mdi-close" @click="isOpen = false" />
      </q-toolbar>
      <div class="scroll-y">
        <q-card-section class="full-width">
          <app-settings-adb />
        </q-card-section>
        <q-card-section class="full-width">
          <app-settings-shell />
        </q-card-section>
        <q-card-section class="full-width">
          <app-settings-known-devices />
        </q-card-section>
      </div>
    </q-card>
  </q-dialog>
</template>
