import type { LocalFileInfo } from '../types/LocalFile'

export namespace FileUtils {

  export const isUploadFileObject = (file: object): file is LocalFileInfo => {
    return Object.hasOwn(file, 'file')
  }

  export type DataTransferResult = {
    name: string
    isFile: boolean
    data: LocalFileInfo[]
  }

  export async function flatDataTransferItems(items?: DataTransferItemList): Promise<DataTransferResult[]> {

    const result: DataTransferResult[] = []

    if (!items) {
      return result
    }

    const itemPromises: Promise<void>[] = []
    for (let i = 0; i < items.length; i++) {
      const item = items[i]
      if (item.kind === 'file') {
        const resultList: LocalFileInfo[] = []
        let entry = item.webkitGetAsEntry()
        if (!entry) {
          continue
        }
        let dataTransferResult: DataTransferResult = {
          name: entry.name,
          isFile: entry.isFile,
          data: resultList
        }
        result.push(dataTransferResult)
        itemPromises.push(getFileFromEntryRecursively(entry, resultList))
      }
    }

    for (let i = 0; i < itemPromises.length; i++) {
      await itemPromises[i]
    }

    return result
  }

  function getFileFromEntryRecursively(entry: FileSystemEntry, resultList: LocalFileInfo[]): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      if (entry.isFile) {
        (entry as FileSystemFileEntry).file(
          file => {
            resultList.push({ file, fullPath: entry.fullPath })
            resolve()
          },
          e => {
            reject()
            console.error(e)
          })
      } else {
        const reader = (entry as FileSystemDirectoryEntry).createReader()
        reader.readEntries(
          entries => {
            let subPromise = entries.map(entry => getFileFromEntryRecursively(entry, resultList))
            Promise.all(subPromise).then(() => resolve())
          },
          e => {
            reject()
            console.error(e)
          }
        )
      }
    })
  }

}
