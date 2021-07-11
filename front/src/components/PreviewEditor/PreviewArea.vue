<template>
  <div class="border border-gray-800 w-full textarea-height ml-1 focus:outline-blue-50 p-4">
    <div v-html="previewMarkdown" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    previewMarkdown: {
      type: String,
      required: true
    }
  }
});
</script>

<style scoped>
.previewarea-height {
  height: 80vh;
}
</style>