<template>
  <textarea
    v-model="markdownText"
    class="border border-gray-800 w-full textarea-height ml-1 focus:outline-blue-50 p-4"
  />
  {{ html }}
  <!-- {{ markdownText }} -->
</template>

<script lang="ts">
import { defineComponent, ref, watchEffect } from 'vue';
import marked from 'marked';

export default defineComponent({
  setup() {
    const markdownText = ref<string>('');
    const html = ref<string>('');

    /** 入力されたmarkdownテキストをHTML形式に変換させる */
    watchEffect(() => {
      html.value = marked(markdownText.value);
    });

    return {
      markdownText,
      html
    };
  }
});
</script>

<style>
.textarea-height {
  height: 80vh;
}
</style>