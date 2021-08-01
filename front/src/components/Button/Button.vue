<template>
  <div>
    <button
      class="border border-gray-300 bg-white shadow-sm appearance-none rounded hover:border-blue-300
			active:bg-gray-300"
    >
      <div class="m-1.5 text-gray-600 text-xs">
        {{ label }}
      </div>
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
	props: {
		label: {
			type: String,
			required: true
		}
	}
});
</script>
