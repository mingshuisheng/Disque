package fileUtils

import (
	"github.com/google/uuid"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

func Init() {
	rand.Seed(time.Now().Unix())
}

func SaveFile(file *multipart.FileHeader) (fileName string, dst string, ext string, err error) {
	fileName = file.Filename
	ext = getExt(file)
	dst = generatePath() + generateFileName(file, ext)
	err = SaveFileToDisk(file, dst)
	if err != nil {
		return "", "", "", err
	}
	return
}

func generatePath() string {
	level1 := rand.Uint32() % 32
	level2 := rand.Uint32() % 32
	level3 := rand.Uint32() % 32
	path := "./upload/file/" + strconv.Itoa(int(level1)) + "/" + strconv.Itoa(int(level2)) + "/" + strconv.Itoa(int(level3))
	os.MkdirAll(path, os.ModePerm)
	return path + "/"
}

func generateFileName(file *multipart.FileHeader, ext string) string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	if ext == "" {
		return id
	}
	return id + "." + ext
}

func getExt(file *multipart.FileHeader) string {
	fileName := file.Filename
	strs := strings.Split(fileName, ".")
	if len(strs) == 1 {
		return ""
	}
	return strs[len(strs)-1]
}

func SaveFileToDisk(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
