import {Component} from "solid-js";
import {ClassNameUtils} from "../../utils/ClassNameUtils";
import "./HomeLeft.scss"

const {rootClass, className} = ClassNameUtils.create("disque-home-left");
export const HomeLeft: Component = () => {
  return(
    <div class={rootClass()}>
    </div>
  )
}
