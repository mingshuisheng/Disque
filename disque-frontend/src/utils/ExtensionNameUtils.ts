export namespace ExtensionNameUtils {

  const fileTypeList = [
    // 图片类型
    {'typeName': 'image', 'types': ['png', 'jpg', 'jpeg', 'bmp', 'gif']},
    // 文本类型
    {'typeName': 'txt', 'types': ['txt']},
    // excel类型
    {'typeName': 'excel', 'types': ['xls', 'xlsx']},
    {'typeName': 'word', 'types': ['doc', 'docx']},
    {'typeName': 'pdf', 'types': ['pdf']},
    {'typeName': 'ppt', 'types': ['ppt']},
    // 视频类型
    {'typeName': 'video', 'types': ['mp4', 'm2v', 'mkv']},
    // 音频
    {'typeName': 'radio', 'types': ['mp3', 'wav', 'wmv']}
  ]

  export function getFileTypeByExtension(ext: string): string {
    if(ext === ''){
      return ''
    }

    for (let i = 0; i < fileTypeList.length; i++) {
      const fileTypeItem = fileTypeList[i]
      const typeName = fileTypeItem.typeName
      const types = fileTypeItem.types
      const result = types.some(function (item) {
        return item === ext;
      });
      if (result) {
        return typeName
      }
    }
    return ''
  }
}
