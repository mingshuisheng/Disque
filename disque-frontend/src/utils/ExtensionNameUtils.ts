export namespace ExtensionNameUtils {

  export type FileType = 'image' | 'txt' | 'excel' | 'word' | 'pdf' | 'ppt' | 'video' | 'radio' | ''
  export type FileTypeItem = {
    typeName: FileType,
    types: string[]
  }

  const fileTypeList: FileTypeItem[] = [
    // 图片类型
    { typeName: 'image', types: ['png', 'jpg', 'jpeg', 'bmp', 'gif', 'svg', 'ico'] },
    // 文本类型
    { typeName: 'txt', types: ['txt'] },
    // excel类型
    { typeName: 'excel', types: ['xls', 'xlsx'] },
    { typeName: 'word', types: ['doc', 'docx'] },
    { typeName: 'pdf', types: ['pdf'] },
    { typeName: 'ppt', types: ['ppt'] },
    // 视频类型
    { typeName: 'video', types: ['mp4', 'm2v', 'mkv'] },
    // 音频
    { typeName: 'radio', types: ['mp3', 'wav', 'wmv'] }
  ]

  export function getFileTypeByExtension(ext: string): FileType {
    if (ext === '') {
      return ''
    }

    for (let i = 0; i < fileTypeList.length; i++) {
      const fileTypeItem = fileTypeList[i]
      const typeName = fileTypeItem.typeName
      const types = fileTypeItem.types
      const result = types.some(function(item) {
        return item === ext
      })
      if (result) {
        return typeName
      }
    }
    return ''
  }

  export function getExtensionByName(fileName: string): string{
    if(!fileName.includes(".")){
      return ""
    }

    return fileName.substring(fileName.lastIndexOf(".") + 1)
  }

  export function getExtensionByFile(file: File): string{
    return getExtensionByName(file.name)
  }

}
