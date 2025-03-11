<script setup lang="ts">
import { copyToClipboard } from "quasar";

const { label, error } = defineProps<{
  label: string;
  error: string;
}>();

const emit = defineEmits<{
  "update:error": [error: string];
}>();
</script>

<template>
  <q-banner v-if="error" class="bg-red text-white" inline-actions>
    <div class="row">
      <div class="col">
        {{ label }}
        <pre class="ellipsis-3-lines" :title="error">{{ error }}</pre>
      </div>
      <div class="column no-wrap">
        <q-btn
          dense
          flat
          round
          icon="mdi-close"
          @click="emit('update:error', '')"
        />
        <q-btn
          dense
          flat
          round
          icon="mdi-content-copy"
          title="Copy to clipboard"
          @click="copyToClipboard(error)"
        />
      </div>
    </div>
  </q-banner>
</template>
