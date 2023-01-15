<template>
  <el-table class='upload-task-table' :data='uploadTasks' style='width: 100%' table-layout='fixed'>
    <el-table-column label='名称'>
      <template #default='{row}'>
        <el-row class='name-column'>
          <img :src='row.files.length > 1? folderIcon: UploadFile.getIconByFile(row.files[0])'
               style='height: 50px; width: 50px;'
               alt=''>
          <span>{{ row.name }}</span>
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
      <template #default='{row}'>
        <span v-if='row.files.length > 1'>{{ row.completedQuantity }}项/{{ row.files.length }}项</span>
        <span v-else>{{row.files[0].size}}b</span>
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
import { UploadFile } from '../../utils/UploadFile'

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
      overflow: hidden;
      border-radius: 50%;

      &:hover {
        cursor: pointer;
        color: var(--el-color-primary-light-3);
        background-color: var(--el-color-primary-light-9);
      }

      &:active {
        cursor: pointer;
        color: var(--el-color-primary);
        background-color: var(--el-color-primary-light-9);
      }
    }
  }
}
</style>
