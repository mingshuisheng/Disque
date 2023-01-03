package main

import (
	"disqueBackend/dao"
	"disqueBackend/router"
	"disqueBackend/utils/fileUtils"
)
import "github.com/gin-gonic/gin"

func main() {
	fileUtils.Init()
	dao.Init()
	engine := router.Init(gin.DebugMode)
	engine.Run(":9000")
}
