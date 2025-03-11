<script setup lang="ts">
import {
  AdbVersion,
  GetAdbPath,
  PickFilePath,
  SetAdbPath,
} from "@/go/main/App";
import { BrowserOpenURL } from "@/runtime";
import { useDebounce } from "@vueuse/core";
import { onBeforeMount, shallowRef, watch } from "vue";

const adbPath = shallowRef("");

async function pickAdbPath() {
  const newPath = await PickFilePath();
  if (!newPath) {
    // No file selected
    return;
  }
  adbPath.value = newPath;
}

onBeforeMount(() => {
  void loadAdbPath();
});

watch(useDebounce(adbPath, 100), (path) => {
  void SetAdbPath(path);
});

async function loadAdbPath() {
  adbPath.value = await GetAdbPath();
}

const adbCheckResult = shallowRef<{
  success: boolean;
  message: string;
} | null>(null);

async function checkAdb() {
  adbCheckResult.value = null;
  try {
    adbCheckResult.value = {
      success: true,
      message: await AdbVersion(),
    };
  } catch (e) {
    adbCheckResult.value = {
      success: false,
      message: e as string,
    };
  }
}

const platformToolsUrl =
  "https://developer.android.com/tools/releases/platform-tools#downloads";
</script>

<template>
  <q-card>
    <q-card-section>
      <q-input
        v-model.lazy="adbPath"
        label="ADB binary path"
        hint="Absolute path or binary name"
        class="q-mb-sm"
      >
        <template #append>
          <q-btn
            flat
            round
            icon="mdi-attachment"
            @click="pickAdbPath"
            title="Pick path"
          />
        </template>
      </q-input>
      <q-btn icon="mdi-check" label="Check" @click="checkAdb" />
    </q-card-section>
    <q-banner
      v-if="adbCheckResult"
      class="text-white full-width overflow-auto"
      :class="adbCheckResult.success ? 'bg-green ' : 'bg-red'"
    >
      <pre>{{ adbCheckResult.message }}</pre>
    </q-banner>
    <q-card-actions align="right">
      <a
        :href="platformToolsUrl"
        :title="platformToolsUrl"
        @click.prevent="BrowserOpenURL(platformToolsUrl)"
      >
        Download Android Platform Tools
      </a>
    </q-card-actions>
  </q-card>
</template>
