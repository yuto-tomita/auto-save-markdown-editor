<template>
  <div class="absolute top-0 left-0 w-full h-full bg-gray-500 opacity-50 inline">
    <img
      src="../../assets/spinning-circles.svg"
      class="absolute top-2/4 left-2/4"
    >
  </div>
  <!-- w-screen h-full opacity-50 bg-gray-500 -->
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
	props: {
		isLoading: {
			type: Boolean,
			required: true
		}
	}
});
</script>