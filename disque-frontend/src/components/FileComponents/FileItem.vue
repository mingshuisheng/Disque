<template>
  <file-base @click="emits('click')" :text="text" :icon="unKnownIcon" />
</template>

<script setup lang="ts">
defineOptions({
  name: "FileItem"
})

const {text} = $defineProps<{
  text: string
}>()

const emits = defineEmits<SE<{
  click(): void
}>>()

</script>
