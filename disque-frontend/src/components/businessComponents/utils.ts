import { ElMessageBox, ElMessage } from 'element-plus'
import 'element-plus/theme-chalk/el-message.css'
import 'element-plus/theme-chalk/el-message-box.css'
import type { FileData } from '../../types'
import { deleteFile } from '../../api'
import { reloadFileState } from '../../states/fileState'

export const deleteFileForComponent = async (file: FileData) => {
  try {
    await ElMessageBox.confirm(`确定删除${file.Name}`, '删除', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteFile(file.ID)
    await reloadFileState()
  } catch (err) {
    console.log(err)
    if (err != 'cancel') {
      ElMessage.error('删除文件失败')
    }
  }
}
