import {JSX, splitProps} from "solid-js";
import {chainHandlers, hope, isFunction, menuTriggerStyles, useStyleConfig} from "@hope-ui/solid";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import {usePopMenuContext} from "./PopMenu";
import classNames from "classnames";


export type PopMenuTriggerProps<C extends  ElementType="div"> = HTMLHopeProps<C>;

const hopeMenuTriggerClass = "hope-menu__trigger";

export function PopMenuTrigger<C extends ElementType ="div">(props: PopMenuTriggerProps<C>) {
  const theme = useStyleConfig().Menu;

  const popMenuContext = usePopMenuContext();

  const [local, others] = splitProps(props as PopMenuTriggerProps<"div">, [
    "ref",
    "class",
    "onContextMenu",
    "onKeyDown",
    "onBlur"
  ])

  const assignTriggerRef = (el: HTMLDivElement) => {
    if(isFunction(local.ref)){
      local.ref(el)
    } else {
      local.ref = el
    }
  }

  const onContextMenu = (event: MouseEvent) => {
    popMenuContext.onTriggerContextMenu(event)
  }

  const onKeyDown: JSX.EventHandlerUnion<HTMLDivElement, KeyboardEvent> = event => {
    chainHandlers(popMenuContext.onTriggerKeyDown, local.onKeyDown)(event)
  }

  const onBlur: JSX.EventHandlerUnion<HTMLDivElement, FocusEvent> = event => {
    chainHandlers(popMenuContext.onTriggerBlur, local.onBlur)(event);
  };

  const classes = () => classNames(local.class, hopeMenuTriggerClass, menuTriggerStyles());


  return(
    <hope.div
      ref={assignTriggerRef}
      id={popMenuContext.state.triggerId}
      aria-haspopup="menu"
      aria-controls={popMenuContext.state.menuContentId}
      aria-expanded={popMenuContext.state.opened}
      class={classes()}
      __baseStyle={theme?.baseStyle?.trigger}
      onContextMenu={onContextMenu}
      onKeyDown={onKeyDown}
      onBlur={onBlur}
      {...others}
    />
  )
}
