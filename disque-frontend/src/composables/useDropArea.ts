import type { GeneralEventListener, MaybeComputedRef } from '@vueuse/core'
import { isClient } from '@vueuse/core'


type DragListener = GeneralEventListener<DragEvent>


export function useDropArea(target: MaybeComputedRef<HTMLElement | null | undefined>, onDrop?: DragListener, globalListen = true) {
  const isOverDropZone = ref(false)
  let counter = 0
  const isGlobalOverDropZone = ref(false)
  let globalCounter = 0

  if (isClient) {

    const dragEnter: DragListener = event => {
      event.preventDefault()
      counter += 1
      isOverDropZone.value = true
    }

    const dragOver: DragListener = event => event.preventDefault()

    const dragLeave: DragListener = event => {
      event.preventDefault()
      counter -= 1
      if (counter === 0) {
        isOverDropZone.value = false
      }
    }

    const dragDrop: DragListener = event => {
      event.preventDefault()
      counter = 0
      isOverDropZone.value = false
      isGlobalOverDropZone.value = false
      onDrop?.(event)
    }

    useEventListener<DragEvent>(target, 'dragenter', dragEnter)
    useEventListener<DragEvent>(target, 'dragover', dragOver)
    useEventListener<DragEvent>(target, 'dragleave', dragLeave)
    useEventListener<DragEvent>(target, 'drop', dragDrop)

    if (globalListen) {
      useEventListener<DragEvent>(document.body, 'dragenter', event => {
        event.preventDefault()
        globalCounter += 1
        isGlobalOverDropZone.value = true
      })
      useEventListener<DragEvent>(document.body, 'dragover', dragOver)
      useEventListener<DragEvent>(document.body, 'dragleave', event => {
        event.preventDefault()
        globalCounter -= 1
        if (globalCounter === 0) {
          isGlobalOverDropZone.value = false
        }
      })
    }

  }

  return {
    isOverDropZone,
    isGlobalOverDropZone
  }
}
