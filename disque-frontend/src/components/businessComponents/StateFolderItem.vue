<template>
  <folder-item :title='file.Name' @click='openFolder'>
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
  </folder-item>
  <rename-dialog $showRenameDialog='showRenameDialog' :file='file'></rename-dialog>
</template>

<script setup lang='ts'>
import type { FileData } from '../../types'
import { DownloadFile } from '../../utils/DownloadFile'
import { deleteFileForComponent } from './utils'

defineOptions({
  name: 'StateFolderItem'
})

const { file } = $defineProps<{
  file: FileData
}>()

let router = useRouter()

const showRenameDialog = ref(false)

const openFolder = () => router.push(`/folder/${file.ID}`)

const download = () => DownloadFile.download(file)

</script>

<style scoped lang='scss'>
.file-name{
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100px;
}
</style>
