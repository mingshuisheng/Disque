<template>
  <div @click="onItemClick" class="context-menu-item">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { useContextShow } from './contextMenuState'

defineOptions({
  name: "ContextMenuItem"
})

const emits = defineEmits<SE<{
  click(): void
}>>()

const {setCurrentShow} = useContextShow()

const onItemClick = () => {
  emits('click')
  setCurrentShow("")
}

</script>

<style scoped lang="scss">
.context-menu-item {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 40px;
  width: 220px;
  overflow: hidden;

  &:hover {
    cursor: pointer;
    background-color: var(--el-menu-hover-bg-color);
  }
}
</style>
