<script setup lang="ts">
import { RefreshKnownDevices } from "@/go/main/App";
import { BrowserOpenURL } from "@/runtime";
import { shallowRef } from "vue";

const refreshing = shallowRef(false);
const error = shallowRef("");

async function refreshList() {
  refreshing.value = true;
  try {
    await RefreshKnownDevices();
  } catch (e) {
    error.value = e as string;
  } finally {
    refreshing.value = false;
  }
}

const devicesListPage =
  "https://storage.googleapis.com/play_public/supported_devices.html";
</script>

<template>
  <q-card>
    <q-card-section class="column">
      <div class="text-h">Known Devices List</div>
    </q-card-section>
    <q-card-section>
      <q-btn
        icon="mdi-refresh"
        label="Refresh"
        :loading="refreshing"
        @click="refreshList"
      />
    </q-card-section>
    <q-banner v-if="error" class="bg-red text-white full-width overflow-auto">
      <pre>{{ error }}</pre>
    </q-banner>
    <q-card-actions align="right">
      <a
        :href="devicesListPage"
        :title="devicesListPage"
        @click.prevent="BrowserOpenURL(devicesListPage)"
      >
        Open Devices list page
      </a>
    </q-card-actions>
  </q-card>
</template>
