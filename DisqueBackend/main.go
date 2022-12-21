package main

import "fmt"
import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/data", func(context *gin.Context) {
		fmt.Println("request data")
		context.JSON(http.StatusOK, gin.H{
			"name": "zhangsan",
			"age":  20,
		})
	})

	engine.Run(":9000")
}
