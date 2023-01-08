<template>
  <div @click="emits('click')" class="context-menu-item">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "ContextMenuItem"
})

const emits = defineEmits<SE<{
  click(): void
}>>()

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
