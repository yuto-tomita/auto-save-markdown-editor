<template>
  <div
    class="w-56 mt-4 h-full border-white shadow-md rounded-md ml-1 cursor-pointer"
    @click="openDraft"
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
import { DraftList } from '@/store/type';

export default defineComponent({
	props: {
		draft: {
			type: Object as PropType<DraftList>,
			required: true
		}
	},
	emits: ['callDraft'],
	setup (props, { emit }) {
		const openDraft = () => {
			emit('callDraft', props.draft.Markdown_text);
		};

		return {
			openDraft
		};
	}
});
</script>

<style>

</style>