import {createSignal} from "solid-js";
import {FileData} from "../../types";
import {loadAllParents, loadFileInfo, loadFileList} from "../../api";

// const [currentPath, setCurrentPath] = createSignal("")

const [fileList, setFileList] = createSignal(new Array<FileData>());

// let PathList: FileData[] = []

const [pathList, setPathList] = createSignal<FileData[]>([])

const rootFile: FileData = {
  ID: 0,
  Name: "",
  ParentID: 0,
  IsDir: true,
  CreatedAt: new Date(),
  UpdatedAt: new Date(),
  DeletedAt: new Date()
}

const [currentFile, setCurrentFile] = createSignal<FileData>(rootFile);

const intFileList = async (ID: number) => {
  if (ID != 0) {
    let fileData = await loadFileInfo(ID);

    let currentPathList = [...pathList()]

    //新开页面
    if (currentPathList.length === 0) {
      currentPathList = await loadAllParents(ID)
    } else {
      //增
      if (currentPathList[currentPathList.length - 1].ID === fileData.ParentID) {
        currentPathList.push(fileData)
      } else if (currentPathList[currentPathList.length - 1].ParentID === fileData.ID) {
        //删
        currentPathList.pop()
        fileData = currentPathList.length !== 0? currentPathList[currentPathList.length - 1]: rootFile
      }
    }

    //初始化
    setCurrentFile(fileData)
    setPathList(currentPathList)
  } else {
    setPathList([])
    setCurrentFile(rootFile)
  }
  const files = await loadFileList(ID);
  setFileList(files)
}

const reloadFileList = async () => {
  const files = await loadFileList(currentFile().ID)
  setFileList(files)
}


export {
  fileList,
  intFileList,
  currentFile,
  reloadFileList,
  pathList
}
