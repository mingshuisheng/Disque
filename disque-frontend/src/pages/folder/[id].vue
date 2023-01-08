<template>
  <el-container class="folder-root">
    <el-header class="folder-header">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/folder/0' }">文件</el-breadcrumb-item>
        <el-breadcrumb-item v-for="parent in fileParents" :to="{path: `/folder/${parent.ID}`}" :key="parent.ID">
          {{ parent.Name }}
        </el-breadcrumb-item>
      </el-breadcrumb>
      <el-row>
        <el-button @click="showNewFolder = true" type="primary">新建文件夹</el-button>
        <el-button type="danger">上传文件</el-button>
      </el-row>
    </el-header>
    <el-main>
      <div class="file-list">
        <state-folder-item v-for="file in fileState" :key="file.ID" :file="file"/>
      </div>
    </el-main>
  </el-container>

  <el-dialog $="showNewFolder" title="新建文件夹">
    <el-form :model="form" ref="ruleFormRef" :rules="rules">
      <el-form-item label="文件夹名称" prop="fileName">
        <el-input maxlength="20" autocomplete="off" v-model="form.fileName"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="showNewFolder = false" type="danger">取消</el-button>
        <el-button type="primary" @click="submitForm">提交</el-button>
      </span>
    </template>
  </el-dialog>

</template>

<script setup lang="ts">
import {currentFile, fileParents, fileState, initFileState, loadNewFile, reloadFileState} from "../../states/fileState";
import {makeDir} from "../../api";
import type {FormInstance, FormRules} from "element-plus";

defineOptions({
  name: "FolderPage"
})
let route = useRoute();

const ruleFormRef = ref<FormInstance>()
const form = reactive({
  fileName: ''
})
const rules = reactive<FormRules>({
  fileName: [{required: true, message: '文件夹名称不能为空', trigger: 'blur'}]
})

const submitForm = async () => {
  if (!ruleFormRef.value) return
  try {
    let valid = await ruleFormRef.value.validate();
    if (!valid) {
      return
    }

    await handlerMakeDir()
  } catch (err) {
    console.log(err)
  }
}

const showNewFolder = ref(false)

const handlerMakeDir = async () => {
  if (form.fileName.trim() === "") {
    return
  }
  await makeDir(currentFile.value.ID, form.fileName)
  await reloadFileState()
  form.fileName = ""
  showNewFolder.value = false
}


watch(() => route.params.id, () => loadNewFile(route.params.id))

onMounted(initFileState)

const fileId = route.params.id
</script>

<style scoped lang="scss">
.folder-root {
  width: 100%;
  height: 100%;

  .folder-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .file-list {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    //height: 100%;
    gap: 30px;
  }
}

.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>
