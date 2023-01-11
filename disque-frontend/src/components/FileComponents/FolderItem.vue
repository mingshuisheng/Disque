<template>
  <ContextMenu>
    <template #default>
      <file-base :title='title' @click="emits('click')" :icon='folderIcon'>
        <slot></slot>
      </file-base>
    </template>
    <template #menu>
      <context-menu-content>
        <slot name='menu'></slot>
      </context-menu-content>
    </template>
  </ContextMenu>
</template>

<script setup lang='ts'>
import {folderIcon} from '../../assets'

defineOptions({
  name: 'FolderItem'
})

const emits = defineEmits<SE<{
  click(): void
}>>()

const { title } = defineProps<{
  title?: string
}>()

</script>

