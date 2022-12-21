import {Component} from "solid-js";
import {ClassNameUtils} from "../../utils/ClassNameUtils";
import {BaseProps} from "../../base/BaseProps";
import "./FlexFill.scss"

const {rootClass, className} = ClassNameUtils.create("disque-flex-fill");
export const FlexFill: Component<BaseProps> =(props) => {
  return(
    <div class={rootClass()}>
      <div class={className("content")}>
        {props.children}
      </div>
    </div>
  )
}
