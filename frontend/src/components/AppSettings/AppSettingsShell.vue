<script setup lang="ts">
import {
  GetTerminalCommand,
  PickFilePath,
  SetTerminalCommand,
} from "@/go/main/App";
import { useDebounce } from "@vueuse/core";
import { onBeforeMount, shallowRef, watch } from "vue";

const terminalCommand = shallowRef("");

async function pickTerminalPath() {
  const newPath = await PickFilePath();
  if (!newPath) {
    // No file selected
    return;
  }
  terminalCommand.value = newPath;
}

onBeforeMount(() => {
  void loadTerminalCommand();
});

watch(useDebounce(terminalCommand, 100), (cmd) => {
  void SetTerminalCommand(cmd);
});

async function loadTerminalCommand() {
  terminalCommand.value = await GetTerminalCommand();
}
</script>

<template>
  <q-card>
    <q-card-section>
      <q-input
        v-model.lazy="terminalCommand"
        label="Terminal command"
        hint="Used to start ADB shell"
        placeholder="With arguments, e.g. `xfce4-terminal -e`"
        class="q-mb-sm"
        bottom-slots
      >
        <template #append>
          <q-btn
            flat
            round
            icon="mdi-attachment"
            @click="pickTerminalPath"
            title="Pick path"
          />
        </template>
      </q-input>
    </q-card-section>
  </q-card>
</template>
