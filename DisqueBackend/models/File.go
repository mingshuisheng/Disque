package models

type File struct {
	Model
	Name        string
	ParentID    uint `gorm:"index"`
	IsDir       bool
	LocalFileID uint `json:"-"`
	ExtType     string
	TreeID      string `gorm:"index"`
}

type FileCache struct {
	Model
	LocalPath  string
	FileNumber uint
	DirNumber  uint
	//单位byte
	Size uint
	//标记是否数据不可用
	Dirty FileCacheDirtySign
}

type FileCacheDirtySign uint

type FileCacheDirtyPosition uint

var (
	FileCacheDirtyEnumLocalPath  FileCacheDirtyPosition = 1
	FileCacheDirtySignFileNumber FileCacheDirtyPosition = 1 << 1
	FileCacheDirtySignDirNumber  FileCacheDirtyPosition = 1 << 2
	FileCacheDirtySignSize       FileCacheDirtyPosition = 1 << 3
)

func setPosition(sign FileCacheDirtySign, pos FileCacheDirtyPosition) FileCacheDirtySign {
	return FileCacheDirtySign(uint(sign) | uint(pos))
}

func unsetPosition(sign FileCacheDirtySign, pos FileCacheDirtyPosition) FileCacheDirtySign {
	var res = uint(sign)
	var setUint = ^uint(pos)
	return FileCacheDirtySign(res & setUint)
}

func (sign FileCacheDirtySign) setLocalPath() FileCacheDirtySign {
	return setPosition(sign, FileCacheDirtyEnumLocalPath)
}

func (sign FileCacheDirtySign) unsetLocalPath() FileCacheDirtySign {
	return unsetPosition(sign, FileCacheDirtyEnumLocalPath)
}

func (sign FileCacheDirtySign) setFileNumber() FileCacheDirtySign {
	return setPosition(sign, FileCacheDirtySignFileNumber)
}

func (sign FileCacheDirtySign) unsetFileNumber() FileCacheDirtySign {
	return unsetPosition(sign, FileCacheDirtySignFileNumber)
}

func (sign FileCacheDirtySign) setDirNumber() FileCacheDirtySign {
	return setPosition(sign, FileCacheDirtySignDirNumber)
}

func (sign FileCacheDirtySign) unsetDirNumber() FileCacheDirtySign {
	return unsetPosition(sign, FileCacheDirtySignDirNumber)
}

func (sign FileCacheDirtySign) setSize() FileCacheDirtySign {
	return setPosition(sign, FileCacheDirtySignSize)
}

func (sign FileCacheDirtySign) unsetSize() FileCacheDirtySign {
	return unsetPosition(sign, FileCacheDirtySignSize)
}
