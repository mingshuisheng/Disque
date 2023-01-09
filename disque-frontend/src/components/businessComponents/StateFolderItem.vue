<template>
  <folder-item @click="openFolder">
    <template #default>
      {{ file?.Name || "[文件名错误]" }}
    </template>
    <template #menu>
      <context-menu-item @click="download">下载</context-menu-item>
      <context-menu-item @click="showRenameDialog = true">重命名</context-menu-item>
      <context-menu-item @click="del">删除</context-menu-item>
    </template>
  </folder-item>
  <rename-dialog $showRenameDialog='showRenameDialog' :file='file' ></rename-dialog>
</template>

<script setup lang="ts">
import type {FileData} from "../../types";
import {DownloadFile} from "../../utils/DownloadFile";

defineOptions({
  name: "StateFolderItem"
})

const {file} = $defineProps<{
  file?: FileData
}>()

let router = useRouter();

const showRenameDialog = ref(false)

const openFolder = () => router.push(`/folder/${file?.ID}`)

const download = () => DownloadFile.download(file)

const del = () => {

}

</script>
