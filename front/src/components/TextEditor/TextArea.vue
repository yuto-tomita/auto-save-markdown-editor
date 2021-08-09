<template>
  <div>
    <input
      v-model="title"
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
import { apiStore } from '@/store/api';

export default defineComponent({
  components: {
    PreviewArea
  },
  emits: ['onChangeTextArea'],
  setup(_, { emit }) {
    const markdownText = ref<string>('');
    const conversionToHtml = ref<string>('');
    const title = ref<string>('');
    const store = apiStore();

    /** 入力されたmarkdownテキストをHTML形式に変換させる */
    watchEffect(() => {
      conversionToHtml.value = marked(markdownText.value);
      emit('onChangeTextArea', conversionToHtml);
    });

    /** タイトルとmarkdownテキストの入力値の変更を検知したときに、下書きとしてDBに保存する */
    watchEffect(async (onInvalidate) => {
      // title無しの場合は下書き保存をしない
      if (!title.value) return;
 
      const params = {
        title: title.value,
        markdown_text: markdownText.value
      };

      try {
        await store.saveMarkdownDraft(params);
      } catch (e) {
        console.log(e);
      }
    });

    return {
      markdownText,
      conversionToHtml,
      title
    };
  }
});
</script>

<style scoped>
.textarea-height {
  height: 70vh;
}
</style>