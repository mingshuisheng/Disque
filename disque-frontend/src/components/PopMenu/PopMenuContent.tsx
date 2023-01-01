import {children, createEffect, createSignal, on, Show, splitProps} from "solid-js";
import {usePopMenuContext} from "./PopMenu";
import {
  Box,
  createClassSelector,
  isFunction,
  menuContentStyles,
  menuTransitionName,
  useStyleConfig
} from "@hope-ui/solid";
import {Portal} from "solid-js/web";
import {Transition} from "solid-transition-group";
import {ClickOutside} from "../ClickOutside/ClickOutside";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import classNames from "classnames";


type PopMenuContentProps<C extends ElementType = "div"> = HTMLHopeProps<C>

const hopeMenuContentClass = "hope-menu__content";

export function PopMenuContent<C extends ElementType = "div">(props: PopMenuContentProps<C>) {
  const theme = useStyleConfig().Menu
  const popMenuContext = usePopMenuContext()
  const [local, others] = splitProps(props as PopMenuContentProps<"div">, ["children", "ref", "class"])

  const [isPortalMounted, setIsPortalMounted] = createSignal(false);

  createEffect(on(
    () => popMenuContext.state.opened,
    () => {
      if (popMenuContext.state.opened) {
        setIsPortalMounted(true)
        // setTimeout(() => setShow(true), 10)
      } else {
        setIsPortalMounted(false)
      }
    }
  ))
  const unmountPortal = () => setIsPortalMounted(false);

  const resolvedChildren = children(() => local.children);

  const assignContentRef = (el: HTMLDivElement) => {
    popMenuContext.assignContentRef(el);
    if (isFunction(local.ref)) {
      local.ref(el);
    } else {
      // eslint-disable-next-line solid/reactivity
      local.ref = el;
    }
  };

  const onClickOutside = (event: Event) => popMenuContext.onContentClickOutside(event.target as HTMLElement)

  const classes = () => classNames(local.class, hopeMenuContentClass, menuContentStyles());

  const transitionName = () => {
    switch (popMenuContext.state.motionPreset) {
      case "scale-top-left":
        return menuTransitionName.scaleTopLeft;
      case "scale-top-right":
        return menuTransitionName.scaleTopRight;
      case "scale-bottom-left":
        return menuTransitionName.scaleBottomLeft;
      case "scale-bottom-right":
        return menuTransitionName.scaleBottomRight;
      case "none":
        return "hope-none";
    }
  };

  return (
    <Show when={isPortalMounted()}>
      <Portal>
        <Transition name={transitionName()} appear onAfterExit={unmountPortal}>
          <Show when={popMenuContext.state.opened}>
            <ClickOutside onClickOutside={onClickOutside}>
              <Box
                role="menu"
                tabindex="-1"
                ref={assignContentRef}
                id={popMenuContext.state.menuContentId}
                aria-activedescendant={popMenuContext.state.activeDescendantId}
                aria-labelledby={popMenuContext.state.triggerId}
                class={classes()}
                __baseStyle={theme?.baseStyle?.content}
                onMouseLeave={popMenuContext.onContentMouseLeave}
                {...others}
              >
                {resolvedChildren()}
              </Box>
            </ClickOutside>
          </Show>
        </Transition>
      </Portal>
    </Show>
  )
}

PopMenuContent.toString = () => createClassSelector(hopeMenuContentClass)

