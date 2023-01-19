import type { UploadTask } from '../types/UploadTask'
import { uploadFile } from '../api'
import { currentFile, reloadFileState } from './fileState'
import type { LocalFileInfo } from '../types/LocalFile'

export const uploadTasks = ref<UploadTask[]>([])

let uploadTaskId = 0

export const addUploadTask = (file: File) => {
  const taskId = uploadTaskId++
  uploadTasks.value.push({
    id: taskId,
    files: [file],
    name: file.name,
    uploadState: 'wait',
    precentage: 0,
    completedQuantity: 0
  })

  uploadFile(file, currentFile.value.ID, progressEvent => {
    let task = uploadTasks.value.find(item => item.id == taskId)
    if (!!task) {
      task.precentage = Math.floor(progressEvent.total ? (progressEvent.loaded / progressEvent.total * 100 | 0) : 0)
      if(task.precentage >= 100){
        reloadFileState().then()
      }
    }
  }).then()
}

export const addUploadTaskDir = (files: LocalFileInfo[], name: string) => {
  const taskId = uploadTaskId++
  uploadTasks.value.push({
    id: taskId,
    files: files,
    name: name,
    uploadState: 'wait',
    precentage: 0,
    completedQuantity: 0
  })
  const totalSize = files.map(item => item.file.size).reduce((pre, current) => pre + current, 0)
  let totalLoaded = 0
  const localFileInfoLoaded = files.map(() => 0)
  files.forEach((localFileInfo, index) => {
    uploadFile(localFileInfo, currentFile.value.ID, progressEvent => {
      let task = uploadTasks.value.find(item => item.id == taskId)
      if (!!task) {
        totalLoaded += progressEvent.loaded - localFileInfoLoaded[index]
        localFileInfoLoaded[index] = progressEvent.loaded
        task.precentage = Math.floor(totalLoaded / totalSize * 100)
        if(task.precentage >= 100){
          reloadFileState().then()
        }
      }
    }).then()
  })
}
