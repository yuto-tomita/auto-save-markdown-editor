<template>
  <div class="flex">
    <textarea
      v-model="markdownText"
      class="border border-gray-800 w-full textarea-height ml-1 focus:outline-blue-50 p-4"
    />
    <PreviewArea />
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
    const html = ref<string>('');

    /** 入力されたmarkdownテキストをHTML形式に変換させる */
    watchEffect(() => {
      html.value = marked(markdownText.value);
      emit('onChangeTextArea', html);
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