<script setup lang="ts">
import { adb } from "@/go/models";
import { computed } from "vue";

const { device } = defineProps<{
  device: adb.Device | null;
}>();

function stateLabel(state: adb.DeviceState) {
  // https://developer.android.com/tools/adb#devicestatus
  // https://stackoverflow.com/a/31232709
  return (
    {
      [adb.DeviceState.Offline]:
        "The device is not connected to ADB or is not responding",
      [adb.DeviceState.Device]: "The device is connected to the ADB server",
      [adb.DeviceState.NoDevice]: "There is no device connected",
      [adb.DeviceState.Recovery]: "The device is in recovery mode",
      [adb.DeviceState.Sideload]: "The device is in sideload mode",
      [adb.DeviceState.Unauthorized]:
        "The ADB host's RSA public key has not been added to the device",
      [adb.DeviceState.NoPermissions]:
        "Insufficient permissions to communicate with the device",
      [adb.DeviceState.Host]: "Unknown state",
    }[state] ?? "Unknown state"
  );
}

const label = computed(() =>
  device ? `${stateLabel(device.state)} [${device.state}]` : null,
);

const icon = computed(() =>
  device
    ? {
        [adb.DeviceState.Offline]: "mdi-cellphone-off",
        [adb.DeviceState.Device]: "mdi-cellphone-check",
        [adb.DeviceState.NoDevice]: "mdi-cellphone",
        [adb.DeviceState.Recovery]: "mdi-cellphone-cog",
        [adb.DeviceState.Sideload]: "mdi-cellphone-arrow-down",
        [adb.DeviceState.Unauthorized]: "mdi-cellphone-key",
        [adb.DeviceState.NoPermissions]: "mdi-cellphone-lock",
        [adb.DeviceState.Host]: "mdi-cellphone",
      }[device.state]
    : "mdi-cellphone",
);

const color = computed(() =>
  device
    ? {
        [adb.DeviceState.Offline]: "amber-7",
        [adb.DeviceState.Device]: "green",
        [adb.DeviceState.NoDevice]: "grey",
        [adb.DeviceState.Recovery]: "blue-7",
        [adb.DeviceState.Sideload]: "cyan",
        [adb.DeviceState.Unauthorized]: "red",
        [adb.DeviceState.NoPermissions]: "pink-4",
        [adb.DeviceState.Host]: "brown-4",
      }[device.state]
    : "grey",
);
</script>

<template>
  <q-icon :name="icon" :color="color" :title="label" />
</template>
