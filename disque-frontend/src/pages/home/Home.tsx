import {Component, createResource, For, onMount} from "solid-js";
import {HomeTop} from "./HomeTop";
import {FileData} from "../../types";
import {fileList, intFileList} from "../../global/state/path";
import {Flex} from "@hope-ui/solid";
import {Folder, PopMenuItem} from "../../components";
import {useMatch, useParams} from "@solidjs/router";
import {gotoFolder} from "../../routes";
import {UnknownFile} from "../../components/FileItem/UnknownFile";
import {coverURL} from "../../api/axios";


type FolderItem = {
  name: string
  onClick: () => void
}

function renderFolder(file: FileData) {
  const handlerClick = async (file: FileData) => {
    gotoFolder(file.ID)
  }

  const items: FolderItem[] = [
    {
      name: "下载",
      onClick: () => {
        alert("正在下载")
      }
    },
    {
      name: "重命名",
      onClick: () => {
        alert("重命名")
      }
    },
    {
      name: "移动",
      onClick: () => {
        alert("文件夹移动")
      }
    }
  ]

  const renderItem = (item: FolderItem) => <PopMenuItem onSelect={item.onClick}>{item.name}</PopMenuItem>


  return (
    <Folder onClick={() => handlerClick(file)} name={file.Name}>
      <For each={items}>
        {renderItem}
      </For>
    </Folder>
  )
}

function renderFile(file: FileData){
  const handlerClick = (file: FileData) => {
    alert("打开文件")
  }

  const items: FolderItem[] = [
    {
      name: "下载",
      onClick: () => {
        window.open(coverURL(`/file/download/${file.ID}`))
      }
    },
    {
      name: "重命名",
      onClick: () => {
        alert("重命名")
      }
    },
    {
      name: "移动",
      onClick: () => {
        alert("文件夹移动")
      }
    }
  ]

  const renderItem = (item: FolderItem) => <PopMenuItem onSelect={item.onClick}>{item.name}</PopMenuItem>

  return (
    <UnknownFile onClick={() => handlerClick(file)} name={file.Name}>
      <For each={items}>
        {renderItem}
      </For>
    </UnknownFile>
  )
}

const Home: Component = () => {
  const params = useParams();
  const match = useMatch(() => "/folder/:fileID");
  const loadFileList = () => intFileList(params.fileID ? parseInt(params.fileID) : 0);
  onMount(loadFileList)
  if (!!match()) {
    createResource(() => params.fileID, loadFileList)
  }

  return (
    <Flex bg="$whiteAlpha1" direction="column" minW="$full">
      <HomeTop/>
      <Flex wrap="wrap" overflowY="auto" gap="20px">
        <For each={fileList()}>
          {(fileData) => fileData.IsDir? renderFolder(fileData): renderFile(fileData)}
        </For>
      </Flex>
    </Flex>
  )
}

export default Home
