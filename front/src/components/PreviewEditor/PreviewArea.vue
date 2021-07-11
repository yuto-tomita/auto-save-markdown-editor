<template>
  <div class="border border-gray-800 previewarea-height w-full mx-0.5">
    <span>PreviewArea</span>
    <!-- <div>{{ previewMarkdown }}</div> -->
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