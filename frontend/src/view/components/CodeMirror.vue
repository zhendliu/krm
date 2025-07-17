<script setup>
import {Codemirror} from 'vue-codemirror'
import {StreamLanguage} from '@codemirror/language'
import {yaml} from '@codemirror/legacy-modes/mode/yaml'
import {ref} from 'vue'
const extensions = [StreamLanguage.define(yaml),]
let code = ref('')
const props = defineProps({
    placeHolderMessage : {
        type: String,
        default: "请输入一些内容"
    },
    height: {
        type: String,
        default: "400px"
    },
    width: {
        type: String,
        default: "100%"
    },
})
</script>

<template>
  <Codemirror 
    v-model="code"
    :placeholder="props.placeHolderMessage"
    :style="{height: props.height, width: props.width, border: '1px dashed grey;'}"
    class="codemirror"
    :autofocus="true"
    :indent-with-tab="true"
    :tab-size="2"
    :extensions="extensions"
  />
</template>

<style scoped>
.codemirror{
  text-align: left;
  overflow-y: scroll;
}
</style>
