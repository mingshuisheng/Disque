import { extIconMap, unKnownIcon } from '../assets'
import { ExtensionNameUtils } from './ExtensionNameUtils'

export namespace UploadFile {

  const defaultIcon = unKnownIcon
  export function getIconByFile(file: File): string {
    const ext = ExtensionNameUtils.getExtensionByFile(file)
    let fileType = ExtensionNameUtils.getFileTypeByExtension(ext)

    if(fileType === "image"){
      return URL.createObjectURL(file)
    }

    return extIconMap.get(ext) || extIconMap.get(fileType) || defaultIcon
  }
}
