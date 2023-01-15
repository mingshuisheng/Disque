import type { UploadTask } from '../types/UploadTask'
import { uploadFile } from '../api'
import { currentFile } from './fileState'

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
    }
  }).then()
}

export const addUploadTaskDir = (files: File[], name: string) => {
  const taskId = uploadTaskId++
  uploadTasks.value.push({
    id: taskId,
    files: files,
    name: name,
    uploadState: 'wait',
    precentage: 0,
    completedQuantity: 0
  })
}
