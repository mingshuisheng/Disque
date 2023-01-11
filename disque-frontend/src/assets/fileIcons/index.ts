import unKnownIcon from "./unknown.svg";
import folderIcon from "./folder.svg"
import {
  bmpIcon,
  compressIcon,
  cssIcon,
  dwgIcon,
  dwtIcon,
  dxfIcon,
  exeIcon,
  gifIcon,
  imgIcon,
  jpgIcon,
  linkIcon,
  movIcon,
  mp4Icon,
  musicIcon,
  pdfIcon,
  pngIcon,
  pptIcon,
  rtfIcon,
  statisticsIcon,
  svgIcon,
  txtIcon,
  wordIcon,
  xlsxIcon
} from './ext'

const extIconMap = new Map<string, string>()
extIconMap.set("bmp", bmpIcon)
extIconMap.set("compress", compressIcon)
extIconMap.set("css", cssIcon)
extIconMap.set("dwg", dwgIcon)
extIconMap.set("dwt", dwtIcon)
extIconMap.set("dxf", dxfIcon)
extIconMap.set("exe", exeIcon)
extIconMap.set("gif", gifIcon)
extIconMap.set("image", imgIcon)
extIconMap.set("jpg", jpgIcon)
extIconMap.set("link", linkIcon)
extIconMap.set("mov", movIcon)
extIconMap.set("video", movIcon)
extIconMap.set("mp4", mp4Icon)
extIconMap.set("radio", musicIcon)
extIconMap.set("pdf", pdfIcon)
extIconMap.set("png", pngIcon)
extIconMap.set("ppt", pptIcon)
extIconMap.set("rtf", rtfIcon)
extIconMap.set("statistics", statisticsIcon)
extIconMap.set("svg", svgIcon)
extIconMap.set("txt", txtIcon)
extIconMap.set("word", wordIcon)
extIconMap.set("xlsx", xlsxIcon)
extIconMap.set("excel", xlsxIcon)
extIconMap.set("unknown", unKnownIcon)
extIconMap.set("", unKnownIcon)

export {
  unKnownIcon,
  folderIcon,
  extIconMap
}
