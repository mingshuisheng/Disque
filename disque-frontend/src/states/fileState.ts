import {loadAllParents, loadFileInfo, loadFileList} from "../api";
import type {FileData} from "../types";

const {state: fileState, isReady, isLoading, execute: execute} = useAsyncState((args) => loadFileList(args.id), [], {
  immediate: false,
  delay: 0,
  resetOnExecute: false
})

const fileParents = ref<FileData[]>([])

const currentFile = ref<FileData>({
  ID: 0,
  Name: "",
  ParentID: 0,
  IsDir: true,
  ExtType: '',
  CreatedAt: new Date(),
  UpdatedAt: new Date(),
})

const reloadFileState = () => execute(0, {id: currentFile.value.ID})

const loadNewFile = async (id: number) => {
  currentFile.value = await loadFileInfo(id)
  await execute(0, {id})
  if(id === 0){
    fileParents.value = []
  }else{
    fileParents.value = await loadAllParents(id)
  }
}

export {
  isLoading,
  currentFile,
  fileParents,
  fileState,
  // initFileState,
  reloadFileState,
  loadNewFile
}
