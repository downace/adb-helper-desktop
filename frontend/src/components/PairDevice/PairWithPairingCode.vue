<script setup lang="ts">
import PairWithPairingCodeServiceItem from "@/components/PairDevice/PairWithPairingCodeServiceItem.vue";
import { useAppStore } from "@/store";
import { computed } from "vue";

const emit = defineEmits<{
  pair: [];
}>();

const store = useAppStore();

const pairingServices = computed(() =>
  [...store.services.values()].filter(
    (svc) => svc.type.name === "_adb-tls-pairing._tcp",
  ),
);
</script>

<template>
  <div class="column no-wrap items-center">
    <ol>
      <li>
        Go to
        <strong>Developer options</strong>
        &rarr;
        <strong>Wireless debugging</strong>
      </li>
      <li>
        Select
        <strong> Pair device with pairing code </strong>
      </li>
      <li>
        <div v-if="pairingServices.length > 0">
          Enter pairing code for your device:
        </div>
        <div v-else>Waiting for devices...</div>
      </li>
    </ol>

    <q-list separator class="self-stretch">
      <pair-with-pairing-code-service-item
        v-for="service in pairingServices"
        :service="service"
        @pair="emit('pair')"
      />
    </q-list>
  </div>
</template>
