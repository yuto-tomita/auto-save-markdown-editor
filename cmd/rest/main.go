package main

import (
	"fmt"
	"myapp/pkg/db"
	"net/http"

	"myapp/pkg/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := db.ConnectDB()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.POST("/saveMarkDown", func(req *gin.Context) {
		fmt.Println(req.PostForm("markdown_text"))
	})

	router.GET("/", func(req *gin.Context) {
		draftList := api.GetAll(db)

		req.JSON(http.StatusOK, gin.H {
			"res": draftList.Value,
		})
	})
	router.Run(":3000")
}

