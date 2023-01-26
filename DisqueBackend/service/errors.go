package service

import "emperror.dev/errors"

// var FileExitedError = errors.Sentinel("文件已存在")

var FileNotExitError = errors.Sentinel("文件不存在")
var ParentFileNotExitError = errors.Sentinel("父文件夹不存在")
var FolderCreateError = errors.Sentinel("文件夹创建失败")
