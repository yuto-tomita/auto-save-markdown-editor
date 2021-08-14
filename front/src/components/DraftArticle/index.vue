<template>
  <div
    class="w-56 mt-4 h-full border-white shadow-md rounded-md ml-1 cursor-pointer"
  >
    <div class="border-b border-gray-300">
      タイトル: <span class="font-bold">{{ draft.Title }}</span>
    </div>
    <div class="mt-4">
      {{ draft.Markdown_text }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { Draft } from '@/store/type';

export default defineComponent({
	props: {
		draft: {
			type: Object as PropType<Draft>,
			required: true
		}
	}
});
</script>

<style>

</style>