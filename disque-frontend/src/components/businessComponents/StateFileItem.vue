<template>
  <file-item :title='file.Name' :icon='icon'>
    <template #default>
      <div class='file-name'>
        {{ file.Name || '[文件名错误]' }}
      </div>
    </template>
    <template #menu>
      <context-menu-item @click='download'>下载</context-menu-item>
      <context-menu-item @click='showRenameDialog = true'>重命名</context-menu-item>
      <context-menu-item @click='deleteFileForComponent(file)'>删除</context-menu-item>
    </template>
  </file-item>
  <rename-dialog $showRenameDialog='showRenameDialog' :file='file'></rename-dialog>
</template>

<script setup lang='ts'>
import type { FileData } from '../../types'
import { DownloadFile } from '../../utils/DownloadFile'
import { deleteFileForComponent } from './utils'
import { extIconMap } from '../../assets'
import { ExtensionNameUtils } from '../../utils/ExtensionNameUtils'

defineOptions({
  name: 'StateFileItem'
})

const { file } = $defineProps<{
  file: FileData
}>()

const showRenameDialog = ref(false)

const download = () => DownloadFile.download(file)

const icon = computed<string>((): string => {
  if (extIconMap.has(file.ExtType)) {
    return extIconMap.get(file.ExtType) || ''
  }

  const typeName = ExtensionNameUtils.getFileTypeByExtension(file.ExtType)
  if (extIconMap.has(typeName)) {
    return extIconMap.get(typeName) || ''
  }

  return ''
})

</script>

<style scoped lang='scss'>
.file-name {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100px;
}
</style>
