<template>
  <div
    v-if="isDisplayModal"
    class="absolute w-full h-full bg-gray-200 bg-opacity-70"
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
import { ApiState } from '@/store/type';
import DraftArticle from '@/components/DraftArticle/index.vue';

export default defineComponent ({
  components: {
    DraftArticle
  },
  props: {
    isDisplayModal: {
      type: Boolean,
      required: true
    },
    draftList: {
      type: Array as PropType<ApiState['DraftList']>,
      required: true
    }
  },
  emits: ['callDraft'],
  setup (_, { emit }) {
    const callDraft = (draft: ApiState['DraftList']) => {
      emit('callDraft', draft);
    };

    return {
      callDraft,
    };
  }
});
</script>

<style scoped>
.left-margin {
  left: 15%; 
}
</style>
