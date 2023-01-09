import {handlerResult, httpClient} from "./axios";
import type {FileData} from "../types";
import type { ResponseMsg } from '../types/ResponseMsg'


export const loadFileList = (parentID: number = 0) => handlerResult<FileData[]>(() => httpClient.get(`/file/list/${parentID}`))

export const makeDir = (parentID: number, name: string) => handlerResult<{ msg: number }>(() => httpClient.post("/file", {
  parentID,
  name
}))

export const loadAllParents = (ID: number) => handlerResult<FileData[]>(() => httpClient.get(`/file/parents/${ID}`))

export const loadFileInfo = (ID: number) => handlerResult<FileData>(() => httpClient.get(`/file/info/${ID}`))

export const renameFile = (ID: number, newFileName: string) => handlerResult<ResponseMsg>(() => httpClient.put("/file/rename", {ID, NewFileName: newFileName}))

export const uploadFile = (file: File, parentID: number) => {
  return handlerResult<{ msg: string }>(
    () => {
      const fd = new FormData()
      fd.append("file", file)
      fd.append("parentID", parentID.toString());
      return httpClient.post("/file/upload", fd, {
        onUploadProgress: progressEvent => {
          let percent = progressEvent.total ? (progressEvent.loaded / progressEvent.total * 100 | 0) : 0
          console.log(`parentID progress:${percent}%`)
        }
      })
    }
  )
}
