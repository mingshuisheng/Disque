<template>
  <div class="context-menu-content">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
defineOptions({
  name: "ContextMenuContent"
})

</script>

<style scoped lang="scss">

.context-menu-content {
  display: flex;
  flex-direction: column;
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-base);
  background-color: var(--el-bg-color);
  overflow: hidden;
}
</style>
