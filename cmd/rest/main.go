package main

import (
	"fmt"
	"myapp/pkg/db"
	"net/http"

	"myapp/pkg/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.ConnectDB()
	fmt.Println(db)

	router.POST("/saveMarkDown", )

	router.GET("/", func(req *gin.Context) {
		str := api.GetAll(db)
		fmt.Println(str)
		req.JSON(http.StatusOK, gin.H {
			"message": str,
		})
	})
	router.Run(":3000")
}

