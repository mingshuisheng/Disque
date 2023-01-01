import {
  Accessor,
  createEffect,
  createSignal,
  createUniqueId,
  JSXElement,
  onMount,
  Show,
  splitProps
} from "solid-js";
import {
  Box,
  hope,
  isFunction, menuItemCommandStyles,
  menuItemIconWrapperStyles,
  menuItemStyles, menuItemTextStyles,
  MenuItemVariants,
  useStyleConfig
} from "@hope-ui/solid";
import classNames from "classnames";
import {MenuItemData} from "./PopMenu.utils";
import {MarginProps} from "@hope-ui/solid/dist/styled-system/props/margin";
import {ElementType, HTMLHopeProps} from "@hope-ui/solid/dist/components/types";
import {usePopMenuContext} from "./PopMenu";

type PopMenuItemOptions = Partial<Omit<MenuItemData, "key">> &
  MenuItemVariants & {
  /**
   * The icon to display next to the menu item text.
   */
  icon?: JSXElement;

  /**
   * The space between the icon and the menu item text.
   */
  iconSpacing?: MarginProps["marginRight"];

  /**
   * Right-aligned label text content, useful for displaying hotkeys.
   */
  command?: string;

  /**
   * The space between the command and the menu item text.
   */
  commandSpacing?: MarginProps["marginLeft"];
};

export type PopMenuItemProps<C extends ElementType = "div"> = HTMLHopeProps<C, PopMenuItemOptions>;

const hopeMenuItemClass = "hope-menu__item";
const hopeMenuItemIconWrapperClass = "hope-menu__item__icon-wrapper";
const hopeMenuItemTextClass = "hope-menu__item__text";
const hopeMenuItemCommandClass = "hope-menu__item__command";

export function PopMenuItem(props: PopMenuItemProps) {
  const key = createUniqueId();

  const theme = useStyleConfig().Menu

  const popMenuContext = usePopMenuContext();

  const [index, setIndex] = createSignal<number>(-1);

  let itemRef: HTMLDivElement | undefined;

  const [local, others] = splitProps(props as PopMenuItemProps<"div">, [
    "ref",
    "class",
    "children",
    "colorScheme",
    "icon",
    "iconSpacing",
    "command",
    "commandSpacing",
    "textValue",
    "disabled",
    "closeOnSelect",
    "onSelect",
    "onClick",
  ]);

  const itemData: Accessor<MenuItemData> = () => ({
    key,
    textValue: local.textValue ?? itemRef?.textContent ?? "",
    disabled: !!local.disabled,
    closeOnSelect:
      local.closeOnSelect != null ? !!local.closeOnSelect : popMenuContext.state.closeOnSelect,
    onSelect: local.onSelect,
  });

  const id = () => `${popMenuContext.state.itemIdPrefix}-${index()}`;

  const isActiveDescendant = () => popMenuContext.isItemActiveDescendant(index());

  const assignItemRef = (el: HTMLDivElement) => {
    itemRef = el;

    if (isFunction(local.ref)) {
      local.ref(el);
    } else {
      // eslint-disable-next-line solid/reactivity
      local.ref = el;
    }
  };

  const onItemClick = (event: MouseEvent) => {
    event.stopPropagation();
    popMenuContext.onItemClick(index());
  };

  const onItemMouseMove = (event: MouseEvent) => {
    if (local.disabled) {
      popMenuContext.onItemMouseMove(-1);
    }

    if (isActiveDescendant() || local.disabled) {
      event.preventDefault();
      event.stopPropagation();
      return;
    }

    popMenuContext.onItemMouseMove(index());
  };

  const classes = () => {
    return classNames(
      local.class,
      hopeMenuItemClass,
      menuItemStyles({
        colorScheme: local.colorScheme,
      })
    );
  };

  const iconWrapperClasses = () => {
    return classNames(hopeMenuItemIconWrapperClass, menuItemIconWrapperStyles());
  };

  const textClasses = () => {
    return classNames(hopeMenuItemTextClass, menuItemTextStyles());
  };

  const commandClasses = () => {
    return classNames(hopeMenuItemCommandClass, menuItemCommandStyles());
  };

  onMount(() => {
    setIndex(popMenuContext.registerItem(itemData()));
  });

  createEffect(() => {
    if (isActiveDescendant() && itemRef) {
      popMenuContext.scrollToItem(itemRef);
    }
  });

  return (
    <Box
      ref={assignItemRef}
      role="menuitem"
      id={id()}
      data-active={isActiveDescendant() ? "" : undefined}
      data-disabled={local.disabled ? "" : undefined}
      data-group
      class={classes()}
      __baseStyle={theme?.baseStyle?.item}
      onClick={onItemClick}
      onMouseMove={onItemMouseMove}
      onMouseDown={popMenuContext.onItemMouseDown}
      {...others}
    >
      <Show when={local.icon}>
        <hope.span
          aria-hidden="true"
          class={iconWrapperClasses()}
          __baseStyle={theme?.baseStyle?.itemIconWrapper}
          mr={local.iconSpacing ?? "0.5rem"}
        >
          {local.icon}
        </hope.span>
      </Show>
      <Show when={local.children}>
        <hope.span class={textClasses()} __baseStyle={theme?.baseStyle?.itemText}>
          {local.children}
        </hope.span>
      </Show>
      <Show when={local.command}>
        <hope.span
          class={commandClasses()}
          __baseStyle={theme?.baseStyle?.itemCommand}
          ml={local.commandSpacing ?? "0.5rem"}
        >
          {local.command}
        </hope.span>
      </Show>
    </Box>
  )
}
