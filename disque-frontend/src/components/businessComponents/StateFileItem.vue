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
import type { WatchStopHandle } from '@vue/runtime-core'
import { coverURL } from '../../api/axios'

defineOptions({
  name: 'StateFileItem'
})

const { file } = $defineProps<{
  file: FileData
}>()

const showRenameDialog = ref(false)

const download = () => DownloadFile.download(file)

const icon = ref('')

const reloadIcon = () => {
  const fileType = ExtensionNameUtils.getFileTypeByExtension(file.ExtType)

  if (fileType === 'image') {
    try {
      // let value = await loadImageUrl(file.ID)
      // console.log("url", value)
      icon.value = coverURL(`/file/accessUrl/${file.ID}`)
      return
    } catch {
      console.warn('load image error')
    }
  }

  icon.value = extIconMap.get(file.ExtType) || extIconMap.get(fileType) || ''
}

let stop: WatchStopHandle

onMounted(() => {
  reloadIcon()
  stop = watch(() => file.ExtType, reloadIcon)
})

onUnmounted(() => stop?.())


</script>

<style scoped lang='scss'>
.file-name {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100px;
}
</style>
