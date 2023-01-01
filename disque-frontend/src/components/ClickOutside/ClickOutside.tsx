import { children, PropsWithChildren } from "solid-js";
import {useClickOutside} from "@hope-ui/solid";

export type ClickOutsideProps = PropsWithChildren<{
  /**
   * Callback invoked when the user click outside.
   */
  onClickOutside: (event: Event) => void;
}>;

export function ClickOutside(props: ClickOutsideProps) {
  const resolvedChildren = children(() => props.children);

  useClickOutside({
    element: () => resolvedChildren() as HTMLElement,
    handler: event => props.onClickOutside(event),
  });

  return resolvedChildren;
}
