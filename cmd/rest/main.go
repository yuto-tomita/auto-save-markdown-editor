package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(req *gin.Context) {
		req.JSON(http.StatusOK, gin.H {
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}