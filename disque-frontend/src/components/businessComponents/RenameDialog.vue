<template>
  <el-dialog $='showRenameDialog' title='重命名'>
    <el-form :model='form' ref='ruleFormRef' :rules='rules'>
      <el-form-item label='文件夹名称' prop='fileName'>
        <el-input maxlength='20' autocomplete='off' v-model='form.fileName' />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class='dialog-footer'>
        <el-button @click='showRenameDialog = false' type='danger'>取消</el-button>
        <el-button type='primary' @click='submitForm'>提交</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang='ts'>
import type { FormInstance, FormRules } from 'element-plus'
import { renameFile } from '../../api'
import { reloadFileState } from '../../states/fileState'
import type { FileData } from '../../types'

const {file} = $defineProps<{
  file: FileData
}>()

defineOptions({
  name: 'RenameDialog'
})

const {showRenameDialog} = defineModel<{
  showRenameDialog: boolean
}>()


const ruleFormRef = ref<FormInstance>()
const form = reactive({
  fileName: file.Name || ""
})
const rules = reactive<FormRules>({
  fileName: [{ required: true, message: '文件夹名称不能为空', trigger: 'blur' }]
})

const submitForm = async () => {
  if (!ruleFormRef.value) return
  try {
    await ruleFormRef.value.validate()
    if(form.fileName.trim() === file.Name){
      return
    }
    await renameFile(file.ID || 0, form.fileName.trim())
    showRenameDialog.value = false
    await reloadFileState()
  } catch (err) {
    // console.log(err)
  }
}

</script>

<style scoped lang='scss'>

</style>
