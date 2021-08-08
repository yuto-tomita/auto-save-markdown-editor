<template>
  <div>
    <input
      type="text"
      class="border border-gray-800 w-1/2 mb-2"
      placeholder="タイトル"
    >
  </div>
  <div class="flex">
    <textarea
      v-model="markdownText"
      class="border border-gray-800 w-full textarea-height ml-1 focus:outline-blue-50 p-4"
    />
    <PreviewArea :preview-markdown="conversionToHtml" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watchEffect } from 'vue';
import PreviewArea from '@/components/PreviewEditor/PreviewArea.vue';
import marked from 'marked';

export default defineComponent({
  components: {
    PreviewArea
  },
  emits: ['onChangeTextArea'],
  setup(_, { emit }) {
    const markdownText = ref<string>('');
    const conversionToHtml = ref<string>('');

    /** 入力されたmarkdownテキストをHTML形式に変換させる */
    watchEffect(() => {
      conversionToHtml.value = marked(markdownText.value);
      emit('onChangeTextArea', conversionToHtml);
    });

    return {
      markdownText,
      conversionToHtml
    };
  }
});
</script>

<style scoped>
.textarea-height {
  height: 80vh;
}
</style>