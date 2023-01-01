import {handlerResult, httpClient} from "./axios";
import {FileData} from "../types";


export namespace FileApi{

}

export const loadFileList = (parentID: number = 0) => handlerResult<FileData[]>(() => httpClient.get(`/file/list/${parentID}`))

export const makeDir = (parentID: number, name: string) => handlerResult<{msg: number}>(() => httpClient.post("/file", {parentID, name}))

export const loadAllParents = (ID: number) => handlerResult<FileData[]>(() => httpClient.get(`/file/parents/${ID}`))

export const loadFileInfo = (ID: number) => handlerResult<FileData>(() => httpClient.get(`/file/info/${ID}`))

export const uploadFile = (file: File) => {
  return handlerResult<{msg: string}>(
    () => {
      const fd = new FormData()
      fd.append("file", file)
      fd.append("parent", "12345");
      // return httpClient.post("/file/upload", fd, {headers: {'Content-type': 'multipart/form-data'}})
      return httpClient.post("/file/upload", fd)
    }
  )
}
