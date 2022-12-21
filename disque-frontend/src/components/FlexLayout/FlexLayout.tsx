import {Component, JSXElement, mergeProps, Show} from "solid-js";
import {BaseProps} from "../../base/BaseProps";
import {ClassNameUtils} from "../../utils/ClassNameUtils";
import {FlexFill} from "../FlexFill/FlexFill";
import "./FlexLayout.scss"
export interface FlexLayoutProps extends BaseProps{
  head?: JSXElement
  tail?: JSXElement
  direction?: "row" | "column"
}


const {rootClass, className} = ClassNameUtils.create("disque-flex-layout");

export const FlexLayout: Component<FlexLayoutProps> = (props) => {
  props = mergeProps({
    direction: "row"
  }, props)
  return (
    <div class={rootClass(props.class, props.direction)}>
      <Show when={props.head}>
        <div class={className("head")}>
          {props.head}
        </div>
      </Show>
      <FlexFill>{props.children}</FlexFill>
      <Show when={props.tail}>
        <div class={className("tail")}>tail</div>
      </Show>
    </div>
  )
}
