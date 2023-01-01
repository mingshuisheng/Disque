import {Component, createSignal, For, Show} from "solid-js";
import {currentFile, pathList, reloadFileList} from "../../global/state/path";
import {makeDir, uploadFile} from "../../api";
import {A} from "@solidjs/router"
import {
  Flex,
  Box,
  Button, createDisclosure,
  FormControl,
  FormLabel, Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay, css, ButtonGroup,
} from "@hope-ui/solid"


export const HomeTop: Component = () => {
  const [fileName, setFileName] = createSignal("")
  const {isOpen, onOpen, onClose} = createDisclosure()

  let onMakeDir = async () => {
    if (fileName().trim() === "") {
      return
    }
    // alert(`save dir ${fileName()}`)
    await makeDir(currentFile().ID, fileName())
    onClose()
    await reloadFileList()
  }

  let handleInput = e => setFileName(e.target.value)
  let handleKeyDown = e => {
    if (e.key === "Enter") {
      onMakeDir().then(null)
    }
  }

  const navStyles = css({
    display: "inline"
  })

  const arrowStyles = css({
    marginLeft: "10px",
    marginRight: "10px"
  })

  const titleStyles = css({
    fontSize: "$lg",
    fontWeight: "$bold",
  })

  let inputEl: HTMLInputElement| undefined
  const assignInputEl = (el: HTMLInputElement) => {
    inputEl = el
  }

  const handlerUploadFile = () => {
    inputEl?.click()
  }

  const handlerFileInputChange = (event: Event & {target: any}) => {
    console.log("input:", event.target.files)
    uploadFile(event.target.files[0])
  }

  return (
    <Flex margin="20px" alignItems="center" justifyContent="space-between">
      <Flex class={titleStyles()}>
        <A class={css({
          opacity: pathList().length !== 0 ? "0.5" : "",
          "&:hover": {
            cursor: pathList().length !== 0 ? "pointer": "default",
            color: pathList().length !== 0 ? "$accent10": ""
          }
        })()} href="/">文件</A>
        <Show when={pathList().length !== 0}>
          <Box class={arrowStyles()}>{">"}</Box>
          <nav class={navStyles()}>
            <For each={pathList()}>
              {
                file => <A class={css({
                  opacity: file.ID !== currentFile().ID ? "0.5" : "",
                  "&:hover": {
                    cursor: file.ID === currentFile().ID ? "default": "pointer",
                    color: file.ID !== currentFile().ID ? "$accent10": ""
                  }
                })()} href={`/folder/${file.ID}`}>{file.Name}/</A>
              }
            </For>
          </nav>
        </Show>
      </Flex>
      <ButtonGroup>
        <Button onClick={onOpen} colorScheme="primary">新建文件夹</Button>
        <Button onClick={handlerUploadFile} colorScheme="danger">上传文件</Button>
      </ButtonGroup>
      <Modal opened={isOpen()} onClose={onClose} initialFocus="#dirName">
        <ModalOverlay/>
        <ModalContent>
          <ModalCloseButton/>
          <ModalHeader>请输入文件夹名称</ModalHeader>
          <ModalBody>
            <FormControl id="dirName">
              <FormLabel>文件夹名称</FormLabel>
              <Input onKeyDown={handleKeyDown} value={fileName()} onInput={handleInput} required={true}
                     placeholder="文件夹名称"/>
            </FormControl>
          </ModalBody>
          <ModalFooter>
            <Button onClick={onMakeDir}>保存</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
        <Input ref={assignInputEl} onChange={handlerFileInputChange} style={{display: "none"}} type="file"/>
    </Flex>
  )
}
