import {onScopeDispose} from "@vue/runtime-core";

const currentShowId = ref('')

export const useContextShow = (onIdChange: (newId: string) => void) => {

  useEventListener(window, "contextmenu", () => {
    currentShowId.value = ""
  })

  const stop = watch(currentShowId, onIdChange)
  onScopeDispose(stop)

  return {
    setCurrentShow: (newId: string) => currentShowId.value = newId,
    stop
  }
}
