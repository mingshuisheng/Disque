import {createContext, createSignal, splitProps, useContext} from "solid-js";
import {Box, menuGroupStyles, useStyleConfig} from "@hope-ui/solid";
import classNames from "classnames";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";

export interface PopMenuGroupContextValue {
  setAriaLabelledBy(id: string): void;
}

const PopMenuGroupContext = createContext<PopMenuGroupContextValue>()

export type PopMenuGroupProps<C extends ElementType = "div"> = HTMLHopeProps<C>

const hopeMenuGroupClass = "hope-menu__group";

export function PopMenuGroup<C extends ElementType = "div">(props: PopMenuGroupProps<C>) {
  const theme = useStyleConfig().Menu

  const [ariaLabelledBy, setAriaLabelledBy] = createSignal<string>()

  const [local, others] = splitProps(props as PopMenuGroupProps<"div">, ["class", "children"]);

  const classes = () => classNames(local.class, hopeMenuGroupClass, menuGroupStyles())

  const context: PopMenuGroupContextValue = {
    setAriaLabelledBy,
  }

  return (
    <PopMenuGroupContext.Provider value={context}>
      <Box
        role="group"
        aria-labelledby={ariaLabelledBy()}
        class={classes()}
        __baseStyle={theme?.baseStyle?.group}
        {...others}
      >
        {local.children}
      </Box>
    </PopMenuGroupContext.Provider>
  )
}


export function usePopMenuGroupContext() {
  return useContext(PopMenuGroupContext)
}
