<template>
  <div @click="emits('click')" class="file-base-root">
    <img class="file-base-icon" :src="icon" alt=""/>
    <div class="file-base-name">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Component } from '@vue/runtime-core'

defineOptions({
  name: "FileBase"
})

const {icon} = $defineProps<{
  icon: string | Component,
}>()

const emits = defineEmits<SE<{
  click(): void
}>>()

</script>

<style scoped lang="scss">
.file-base-root {
  display: flex;
  flex-direction: column;
  align-items: center;

  &:hover {
    cursor: pointer;
  }

  .file-base-icon {
    width: 100px;
    height: 100px;
  }
  .file-base-name{

  }
}
</style>
