<template>
  <Modal
    :is-display-modal="isDisplayModal"
    :draft-list="draftList"
  />
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
  <Loading :is-loading="isLoading" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import TextEditor from '@/components/TextEditor/TextArea.vue';
import { apiStore } from '@/store/api';
import { ApiState } from '@/store/type';
import Button from '@/components/Button/Button.vue';
import Loading from '@/components/Loading/index.vue';
import Modal from '@/components/Modal/index.vue';

export default defineComponent({
  components: {
    TextEditor,
    Button,
    Loading,
    Modal
  },
  setup () {
    const store = apiStore();
    const draftList = ref<ApiState['DraftList']>([]);
    const isLoading = ref<boolean>(false);
    const isDisplayModal = ref<boolean>(false);

    const getDraft = async () => {
      isLoading.value = true;

      try {
        await store.getDraftList();
        draftList.value = store.returnCalendarList;

        isDisplayModal.value = true;
      } catch (e) {
        console.log(e);
      }

      isLoading.value = false;
    };

    return {
      store,
      draftList,
      isLoading,
      getDraft,
      isDisplayModal
    };
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
