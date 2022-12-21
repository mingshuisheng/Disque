import {Component} from "solid-js";
import {FileBase, FileBaseProps} from "../FileBase/FileBase";
import {BaseProps} from "../../../base/BaseProps";
import {FolderIcon} from "../../../icons";


export interface FolderProps extends BaseProps {
  name?: string
}

export const Folder: Component<FolderProps> = (props) => {
  return(
    <FileBase icon={<FolderIcon/>} name={props.name}/>
  )
}
