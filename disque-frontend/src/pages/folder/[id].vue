<template>
  <el-container v-loading='isLoading' class='folder-root'>
    <el-header class='folder-header'>
      <el-breadcrumb separator='/'>
        <el-breadcrumb-item :to="{ path: '/folder/0' }">文件</el-breadcrumb-item>
        <el-breadcrumb-item v-for='parent in fileParents' :to='{path: `/folder/${parent.ID}`}' :key='parent.ID'>
          {{ parent.Name }}
        </el-breadcrumb-item>
      </el-breadcrumb>
      <el-row>
        <el-button @click='showNewFolder = true' type='primary'>新建文件夹</el-button>
        <el-upload :before-upload='handlerUploadFile' :show-file-list='false' class='file-upload-btn'>
          <el-button type='danger'>上传文件</el-button>
        </el-upload>
      </el-row>
    </el-header>
    <el-main class='file-list-main' ref='dropZoneRef'>
      <div v-auto-animate class='file-list'>
        <template v-for='file in fileState' :key='file.ID'>
          <state-folder-item v-if='file.IsDir' :file='file' />
          <state-file-item v-if='!file.IsDir' :file='file' />
        </template>
      </div>
      <div :class='{"drag-upload-area": true, "is-over-drag": isOverDropZone}' v-if='isGlobalOverDropZone || isOverDropZone'>
        <div class='upload-contain'>
          <i-ep-upload-filled class='upload-icon' />
          <div class='upload-tips'>文件拖拽到此处</div>
        </div>
      </div>
    </el-main>
  </el-container>

  <el-dialog $='showNewFolder' title='新建文件夹'>
    <el-form :model='form' ref='ruleFormRef' :rules='rules'>
      <el-form-item label='文件夹名称' prop='fileName'>
        <el-input maxlength='20' autocomplete='off' v-model='form.fileName' />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class='dialog-footer'>
        <el-button @click='showNewFolder = false' type='danger'>取消</el-button>
        <el-button type='primary' @click='submitForm'>提交</el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup lang='ts'>
import {
  currentFile,
  fileParents,
  fileState,
  initFileState, isLoading,
  loadNewFile,
  reloadFileState
} from '../../states/fileState'
import { makeDir, uploadFile } from '../../api'
import type { FormInstance, FormRules, UploadRawFile } from 'element-plus'
import { addUploadTask, addUploadTaskDir } from '../../states/uploadState'
import { useDropArea } from '../../composables'
import { FileUtils } from '../../utils/FileUtils'

defineOptions({
  name: 'FolderPage'
})
let route = useRoute()

const ruleFormRef = ref<FormInstance>()
const form = reactive({
  fileName: ''
})
const rules = reactive<FormRules>({
  fileName: [{ required: true, message: '文件夹名称不能为空', trigger: 'blur' }]
})

const submitForm = async () => {
  if (!ruleFormRef.value) return
  try {
    await ruleFormRef.value.validate()
    await handlerMakeDir()
  } catch (err) {
    // console.log(err)
  }
}

const showNewFolder = ref(false)

const handlerMakeDir = async () => {
  await makeDir(currentFile.value.ID, form.fileName)
  await reloadFileState()
  form.fileName = ''
  showNewFolder.value = false
}

const handlerUploadFile = (rawFile: UploadRawFile) => {
  addUploadTask(rawFile)
  return false
}


watch(() => (route.params as any).id, () => loadNewFile((route.params as any).id))
onMounted(() => loadNewFile((route.params as any).id))



const dropZoneRef = ref<HTMLElement>()

const { isOverDropZone, isGlobalOverDropZone } = useDropArea(dropZoneRef, async event => {
  let results = await FileUtils.flatDataTransferItems(event.dataTransfer?.items)
  for (let i = 0; i < results.length; i++) {
    console.log(`lines ${i}:`, results[i].name)
    if(results[i].isFile){
      addUploadTask(results[i].data[0].file)
      console.log("\t|---", results[i].data[0].path)
    }else {
      addUploadTaskDir(results[i].data.map(item => item.file), results[i].name)
      for (let j = 0; j < results[i].data.length; j++) {
        console.log("\t|---", results[i].data[j].path)
      }
    }
  }
})


</script>

<style scoped lang='scss'>
.folder-root {
  width: 100%;
  height: 100%;

  .folder-header {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .file-upload-btn {
      margin-left: 12px;
    }
  }

  .file-list-main {
    position: relative;

    .file-list {
      display: flex;
      flex-wrap: wrap;
      width: 100%;
      //height: 100%;
      gap: 30px;
    }

    .drag-upload-area {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      position: absolute;
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      background-color: var(--el-fill-color-blank);

      border: 1px dashed var(--el-border-color);
      opacity: 0.5;

      &.is-over-drag {
        border: 2px dashed var(--el-color-primary);
        background-color: var(--el-color-primary-light-9);
      }

      &:hover {
        border-color: var(--el-color-primary);
      }

      .upload-contain {
        transform: translate(-50%, -50%);
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

        .upload-icon {
          $size: 100px;
          width: $size;
          height: $size;
        }


        .upload-tips {
          font-size: 20px;
        }
      }
    }

  }
}

.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
