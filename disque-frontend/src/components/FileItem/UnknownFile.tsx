import {Component, splitProps} from "solid-js";
import {FileBase} from "./FileBase";
import {UnKnownIcon} from "../../icons";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import {PopMenu, PopMenuContent, PopMenuGroup, PopMenuTrigger} from "../PopMenu";

export interface UnknownFilePropsOptions {
  name?: string
}

type UnknownFileProps<C extends ElementType = "div"> = HTMLHopeProps<C, UnknownFilePropsOptions>

type TriggerItemProps<C extends ElementType = "div"> = HTMLHopeProps<C>

export function UnknownFile(props: UnknownFileProps) {

  function TriggerItem<C extends ElementType = "div">(triggerProps: TriggerItemProps<C>) {
    const [_local, others] = splitProps(triggerProps as TriggerItemProps<"div">, ["onClick"])
    return (
      <FileBase {...others} onClick={props.onClick} icon={<UnKnownIcon/>} name={props.name}/>
    )
  }

  return (
    <PopMenu>
      <PopMenuTrigger as={TriggerItem}/>
      <PopMenuContent>
        <PopMenuGroup>
          {props.children}
        </PopMenuGroup>
      </PopMenuContent>
    </PopMenu>
  )
}
