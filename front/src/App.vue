<template>
  <h1 class="text-center">
    AutoSaveMarkdownEditor
  </h1>
  {{ draftList }}
  <div class="m-3">
    <Button
      label="下書きを呼び出す"
      @click="getDraft"
    />
  </div>
  <TextEditor @onChangeTextArea="handleEvent" />
  <div>
    <button @click="saveMarkdown">
      clickMe!
    </button>
    <Loading :is-loading="isLoading" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import TextEditor from '@/components/TextEditor/TextArea.vue';
import { apiStore } from '@/store/api';
import Button from '@/components/Button/Button.vue';
import Loading from '@/components/Loading/index.vue';

export default defineComponent({
  components: {
    TextEditor,
    Button,
    Loading
  },
  setup () {
    const store = apiStore();
    const draftList = ref({});
    const isLoading = ref(false);

    const getDraft = async () => {
      isLoading.value = true;

      try {
        await store.getDraftList();
        draftList.value = store.returnCalendarList;
      } catch (e) {
        console.log(e);
      }

      isLoading.value = false;
    };

    async function saveMarkdown () {
      try {
        await store.saveMarkdown();
      } catch (e) {
        console.log(e);
      }
    }

    return { store, draftList, saveMarkdown, isLoading, getDraft };
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
