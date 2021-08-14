<template>
  <div
    class="absolute w-full h-full bg-gray-200 bg-opacity-70"
    @click="emitCloseModal"
  >
    <div class="w-2/3 h-1/2 absolute bg-white top-1/4 left-margin rounded-md border-white overflow-scroll">
      <p>呼び出したい下書きを選択してください</p>
      <div class="flex h-2/3">
        <div
          v-for="(draft, index) in draftList"
          :key="index"
        >
          <DraftArticle
            :draft="draft"
            @click="callDraft(draft)"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { Draft } from '@/store/type';
import DraftArticle from '@/components/DraftArticle/index.vue';

export default defineComponent ({
  components: {
    DraftArticle
  },
  props: {
    draftList: {
      type: Array as PropType<Draft[]>,
      required: true
    }
  },
  emits: ['closeModal'],
  setup (_, { emit }) {
    const emitCloseModal = (draft: Draft) => {
      emit('closeModal', draft);
    };

    return {
      emitCloseModal,
    };
  }
});
</script>

<style scoped>
.left-margin {
  left: 15%; 
}
</style>
