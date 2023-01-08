<template>
  <file-item @click="openFolder">
    <template #default>
      {{ file?.Name || "文件1" }}
    </template>
    <template #menu>
      <context-menu-item @click="download">下载</context-menu-item>
      <context-menu-item @click="rename">重命名</context-menu-item>
      <context-menu-item @click="del">删除</context-menu-item>
    </template>
  </file-item>
</template>

<script setup lang='ts'>
import type { FileData } from '../../types'
import { DownloadFile } from '../../utils/DownloadFile'

defineOptions({
  name: 'StateFileItem'
})

const {file} = $defineProps<{
  file?: FileData
}>()

const download = () => DownloadFile.download(file)

</script>

<style scoped lang='scss'>

</style>
