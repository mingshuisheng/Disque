import {Component, JSXElement} from "solid-js";
import {ClassNameUtils} from "../../../utils/ClassNameUtils";
import {BaseProps} from "../../../base/BaseProps";
import "./FileBase.scss"

export interface FileBaseProps extends Omit<BaseProps, "children"> {
  icon: JSXElement
  name?: JSXElement

  onClick?(): void
}

const {rootClass, className} = ClassNameUtils.create("disque-file-base");
export const FileBase: Component<FileBaseProps> = (props) => {
  return (
    <div onClick={props.onClick} class={rootClass(props.class)}>
      <div class={className("icon")}>{props.icon}</div>
      <div class={className("name")}>{props.name}</div>
    </div>
  )
}
