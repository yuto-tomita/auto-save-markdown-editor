<template>
  <textarea
    v-model="markdownText"
    class="border border-gray-800 w-full textarea-height ml-1 focus:outline-blue-50 p-4"
  />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  setup() {
    const markdownText = ref<string>('');
    return {
      markdownText
    };
  }
});
</script>

<style>
.textarea-height {
  height: 80vh;
}
</style>