import {Component} from "solid-js";
import {ClassNameUtils} from "../../utils/ClassNameUtils";
import {currentPath} from "../../global/state/path";
import "./HomeTop.scss"

const {rootClass, className} = ClassNameUtils.create("disque-home-top");
export const HomeTop: Component = () => {
  return(
    <div class={rootClass()}>
      <div class={className("file-path")}>
        {`文件 > ${currentPath()}`}
      </div>
    </div>
  )
}
