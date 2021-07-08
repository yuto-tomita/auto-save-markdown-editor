<template>
  <textarea
    v-model="inputText"
    class="border border-gray-800 outline-none w-full textarea-height ml-1"
  />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  setup() {
    const inputText = ref<string>('');
    return {
      inputText
    };
  }
});
</script>

<style>
.textarea-height {
  height: 80vh;
}
</style>