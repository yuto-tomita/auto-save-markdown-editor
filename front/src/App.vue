<template>
  <h1 class="text-center">
    AutoSaveMarkdownEditor
  </h1>
  <TextEditor @onChangeTextArea="handleEvent" />
  <button @click="saveMarkdown">
    clickMe!
  </button>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import TextEditor from '@/components/TextEditor/TextArea.vue';
import { apiStore } from '@/store/api';

export default defineComponent({
  components: {
    TextEditor
  },
  setup () {
    const test = apiStore();
    const res = ref({});

    getDraft();

    async function getDraft () {
      try {
        await test.getDraftList();
        res.value = test.returnCalendarList;
        // console.log(test.returnCalendarList);
      } catch (e) {
        console.log(e);
      }
    }

    async function saveMarkdown () {
      try {
        await test.saveMarkdown();
      } catch (e) {
        console.log(e);
      }
    }

    return { test, res, saveMarkdown };
  }
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
