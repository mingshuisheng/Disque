import {Component, JSXElement, splitProps} from "solid-js";
import {FileBase} from "./FileBase";
import {FolderIcon} from "../../icons";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import {PopMenu, PopMenuContent, PopMenuGroup, PopMenuItem, PopMenuTrigger} from "../PopMenu";
import {FileData} from "../../types";

export interface FolderPropsOptions{
  name?: string
}

type FolderProps<C extends ElementType = "div"> = HTMLHopeProps<C, FolderPropsOptions>

type TriggerItemProps<C extends ElementType="div"> = HTMLHopeProps<C>

export const Folder: Component<FolderProps> = (props) => {

  function TriggerItem<C extends ElementType="div" > (triggerProps: TriggerItemProps<C>) {
    const [_local, others] = splitProps(triggerProps as TriggerItemProps<"div">, ["onClick"])
    return (
      <FileBase {...others} onClick={props.onClick} icon={<FolderIcon/>} name={props.name}/>
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
