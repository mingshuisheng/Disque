package logic

type FileExitedError struct {
}

func (err *FileExitedError) Error() string {
	return "file exited error"
}

type FileNotExitError struct {
}

func (err *FileNotExitError) Error() string {
	return "file not exit error"
}

type ParentFileNotExitError struct{}

func (err *ParentFileNotExitError) Error() string {
	return "parent file not exit error"
}

type FolderCreateError struct{}

func (err *FolderCreateError) Error() string {
	return "folder create error"
}
