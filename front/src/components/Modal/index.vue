<template>
  <div
    v-if="isDisplayModal"
    class="absolute w-full h-full bg-gray-200 bg-opacity-70"
  >
    <div class="w-2/3 h-1/2 absolute bg-white top-1/4 left-margin rounded-md border-white">
      <p>呼び出したい下書きを選択してください</p>
      {{ draftList }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { ApiState } from '@/store/type';

export default defineComponent ({
  props: {
    isDisplayModal: {
      type: Boolean,
      required: true
    },
    draftList: {
      type: Array as PropType<ApiState['DraftList']>,
      required: true
    }
  }
});
</script>

<style scoped>
.left-margin {
  left: 15%; 
}
</style>
