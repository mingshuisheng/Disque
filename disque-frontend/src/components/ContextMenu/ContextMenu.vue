<template>
  <div @contextmenu="contextMenu" class='context-menu-root'>
    <slot></slot>
  </div>
  <teleport to="body">
    <OnClickOutside @trigger="closeMenu">
      <div v-if="show" class="context-menu-content-root">
        <slot name="menu"></slot>
      </div>
    </OnClickOutside>
  </teleport>
</template>

<script setup lang="ts">

import {useContextShow} from "./contextMenuState";
const show = ref(false)

const closeMenu = () => {
  show.value = false
}

const instance = getCurrentInstance()
const {setCurrentShow} = useContextShow(newId => {
  if (newId !== instance?.uid.toString()) {
    closeMenu()
  }
})

defineOptions({
  name: "ContextMenu"
})

const x = ref("0px")
const y = ref("0px")

const contextMenu = (event: MouseEvent) => {
  event.stopPropagation()
  event.preventDefault()
  x.value = event.clientX + "px"
  y.value = event.clientY + "px"
  show.value = true
  setCurrentShow(instance?.uid.toString() || "")
}


</script>

<style scoped lang="scss">
.context-menu-root{
  //-webkit-user-drag: none;
}
.context-menu-content-root {
  position: absolute;
  top: v-bind(y);
  left: v-bind(x);
}
</style>
