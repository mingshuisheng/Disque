import {Box, menuLabelStyles, useStyleConfig} from "@hope-ui/solid";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import {createUniqueId, onMount, splitProps} from "solid-js";
import {usePopMenuContext} from "./PopMenu";
import {usePopMenuGroupContext} from "./PopMenuGroup";
import classNames from "classnames";

export type PopMenuLabelProps<C extends ElementType = "div"> = HTMLHopeProps<C>;

const hopeMenuLabelClass = "hope-menu__label";

export function PopMenuLabel(props: PopMenuLabelProps) {
  const defaultIdSuffix = createUniqueId();

  const theme = useStyleConfig().Menu;

  const popMenuContext = usePopMenuContext();
  const popMenuGroupContext = usePopMenuGroupContext();

  const [local, others] = splitProps(props as PopMenuLabelProps<"div">, ["class", "id"]);

  const id = () => local.id ?? `${popMenuContext.state.labelIdPrefix}-${defaultIdSuffix}`;

  const classes = () => classNames(local.class, hopeMenuLabelClass, menuLabelStyles());

  onMount(() => {
    popMenuGroupContext?.setAriaLabelledBy(id());
  });

  return (
    <Box
      id={id()}
      class={classes()}
      __baseStyle={theme?.baseStyle?.label}
      {...others}
    />
  )
}
