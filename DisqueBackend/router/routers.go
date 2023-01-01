package router

import (
	"disqueBackend/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/file/list/:parentID", controller.ListFile)
	r.POST("/file", controller.MakeDir)
	r.GET("/file/parents/:ID", controller.ListAllParents)
	r.GET("/file/info/:ID", controller.FileInfo)
	r.POST("/file/upload", controller.UploadFile)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
