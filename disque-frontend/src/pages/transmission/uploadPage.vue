<template>
  <el-container class='upload-root'>
    <el-header class='upload-header'>
      <span>上传列表 已上传{{ totalPercent }}%</span>
      <el-row>
        <el-button :icon='VideoPlay'>全部开始</el-button>
        <el-button :icon='VideoPause'>全部暂停</el-button>
        <el-button :icon='CircleClose'>全部取消</el-button>
      </el-row>
    </el-header>
    <el-divider style='margin: 0' />
    <el-main>
      <upload-task-table :uploadTasks='tableData'/>
    </el-main>
  </el-container>
</template>

<script setup lang='ts'>
import { VideoPlay, CircleClose, VideoPause } from '@element-plus/icons-vue'
import type { UploadTask } from '../../types/UploadTask'
import UploadTaskTable from '../../components/uploadPage/UploadTaskTable.vue'

defineOptions({
  name: 'UploadPage'
})

const totalPercent = ref(60)

const tableData: UploadTask[] = [
  {
    file: null,
    name: 'k3s',
    precentage: 10,
    uploadState: 'uploading',
    completedQuantity: 0
  },
  {
    file: null,
    name: 'abcd',
    precentage: 60,
    uploadState: 'uploading',
    completedQuantity: 0
  },
  {
    file: null,
    name: '123',
    precentage: 50,
    uploadState: 'uploading',
    completedQuantity: 0
  }
]

</script>

<style scoped lang='scss'>
.upload-root {
  .upload-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
}
</style>
