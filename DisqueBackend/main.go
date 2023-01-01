package main

import (
	"disqueBackend/dao"
	"disqueBackend/router"
)
import "github.com/gin-gonic/gin"

func main() {
	dao.Init()
	engine := router.Init(gin.DebugMode)
	engine.Run(":9000")
}
