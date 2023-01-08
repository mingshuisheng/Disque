<template>
  <ContextMenu>
    <template #default>
      <file-base @click="emits('click')" :icon="unKnownIcon">
        <slot></slot>
      </file-base>
    </template>
    <template #menu>
      <context-menu-content>
        <slot name="menu"></slot>
      </context-menu-content>
    </template>
  </ContextMenu>
</template>

<script setup lang="ts">
defineOptions({
  name: "FileItem"
})

const emits = defineEmits<SE<{
  click(): void
}>>()

</script>
