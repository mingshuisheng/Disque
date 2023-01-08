import type {FileData} from "../types";
import {coverURL} from "../api/axios";

export namespace DownloadFile {
  export const download = (file?: FileData) => {
    if (!file){
      return
    }
    let link = document.createElement('a');
    link.style.display = 'none'
    link.href = coverURL(`/file/download/${file.ID}`)
    link.setAttribute('download',file.Name + (file.IsDir? ".zip" : ""))
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }
}
