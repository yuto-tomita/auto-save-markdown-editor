<template>
  <textarea
    v-model="inputText"
    class="block border-solid border-2 outline-none w-full textarea-height"
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