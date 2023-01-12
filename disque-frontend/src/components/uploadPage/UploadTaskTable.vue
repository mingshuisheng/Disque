<template>
  <el-table class='upload-task-table' :data='uploadTasks' style='width: 100%' table-layout='fixed'>
    <el-table-column label='名称'>
      <template #default='scope'>
        <el-row class='name-column'>
          <img :src='folderIcon' style='height: 50px; width: 50px;' alt=''>
          <span>{{ scope.row.name }}</span>
        </el-row>
      </template>
    </el-table-column>

    <el-table-column align='right'>
      <template #default='scope'>
        <el-row class='control-button'>
          <i-ep-video-pause />
          <i-ep-video-play />
          <i-ep-circle-close />
        </el-row>
      </template>
    </el-table-column>
    <el-table-column label='大小'>
      <template #default='scope'>
        <span>{{ scope.row.completedQuantity }}项/{{ 5 }}项</span>
      </template>
    </el-table-column>
    <el-table-column label='状态'>
      <template #default='scope'>
        <el-progress :percentage='scope.row.precentage' />
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang='ts'>
import type { UploadTask } from '../../types/UploadTask'
import { folderIcon } from '../../assets'

defineOptions({
  name: 'UploadTaskTable'
})

const { uploadTasks } = $defineProps<{
  uploadTasks: UploadTask[]
}>()

</script>

<style scoped lang='scss'>
.upload-task-table {
  .name-column {
    align-items: center;
  }

  .control-button {
    justify-content: flex-end;
    gap: 10px;

    svg {
      width: 20px;
      height: 20px;

      &:hover {
        cursor: pointer;
      }

      &:active {
        opacity: 0.5;
      }
    }
  }
}
</style>
