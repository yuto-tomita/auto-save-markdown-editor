<template>
  <h1>AutoSaveMarkdownEditor</h1>
  <TextEditor @onChangeTextArea="handleEvent" />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import TextEditor from '@/components/TextEditor/TextArea.vue';

export default defineComponent({
  components: {
    TextEditor
  }
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
