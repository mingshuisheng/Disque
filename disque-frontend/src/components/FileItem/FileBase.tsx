import {Component, JSX, JSXElement, splitProps} from "solid-js";
import {Box, Flex, css} from "@hope-ui/solid";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";

export interface FileBasePropsOptions {
  icon: JSXElement
  name?: JSXElement
}

type FileBaseProps<C extends ElementType ="div"> = HTMLHopeProps<C, FileBasePropsOptions>

// const {rootClass, className} = ClassNameUtils.create("disque-file-base");
export const FileBase: Component<FileBaseProps> = (props) => {

  const styles = css({
    borderRadius: "$sm",
    overflow: "hidden",
    "&:hover":{
      // opacity: 0.8,
      cursor: props.onClick? "pointer": "",
      backgroundColor: props.onClick? "$info4": "",
    }
  })

  const [local, others] = splitProps(props, ["name", "icon"])

  // const {name, icon, ...restProps} = props

  return (
    <Flex {...others} class={styles()} direction="column" alignItems="center">
      <Box width="100px" height="100px">{local.icon}</Box>
      <Box>{local.name}</Box>
    </Flex>
  )
}
