package controller

import (
	"disqueBackend/logic"
	"disqueBackend/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

func FileInfo(ctx *gin.Context) {
	IDStr := ctx.Param("ID")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		ID = 0
	}

	if ID == 0 {
		ctx.JSON(http.StatusOK, models.File{})
		return
	}

	file, err := logic.GetFileInfo(ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.File{})
		return
	}
	ctx.JSON(http.StatusOK, file)
}

func ListFile(ctx *gin.Context) {
	parentIDStr := ctx.Param("parentID")
	parentID, err := strconv.ParseUint(parentIDStr, 10, 64)
	if err != nil {
		parentID = 0
	}
	files := logic.GetFileList(parentID)
	ctx.JSON(http.StatusOK, files)
}

type MakeDirParam struct {
	Name     string `json:"name"`
	ParentID uint64 `json:"parentID"`
}

func MakeDir(ctx *gin.Context) {
	param := MakeDirParam{}
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": 404})
		return
	}
	err = logic.MakeDir(param.ParentID, param.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": 500})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": 200})
}

func ListAllParents(ctx *gin.Context) {
	IDStr := ctx.Param("ID")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		ID = 0
	}
	files := logic.GetAllParentFileList(ID)
	ctx.JSON(http.StatusOK, files)
}

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("文件上传失败")
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	parent := ctx.PostForm("parent")

	log.Println("parent", parent)

	dir := "./upload/file"

	os.MkdirAll(dir, os.ModePerm)

	dst := dir + "/" + file.Filename

	ctx.SaveUploadedFile(file, dst)
	ctx.JSON(http.StatusOK, gin.H{"msg": "文件上传成功"})
}
