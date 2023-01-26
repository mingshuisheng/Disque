package controller

import (
	"disqueBackend/service"
	"github.com/gin-gonic/gin"
)

func getFileService(ctx *gin.Context) *service.FileService {
	return service.CreateFileService(ctx.Request.Context())
}
