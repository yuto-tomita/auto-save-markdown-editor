package main

import (
	"fmt"
	"myapp/pkg/db"
	"net/http"

	"myapp/pkg/api"
	"myapp/pkg/model"

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
		var json model.PostJsonForm
		
		if err := req.ShouldBindJSON(&json); err != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(json.MarkdownText)

		req.JSON(http.StatusOK, gin.H{"str": json.MarkdownText})
	})

	router.GET("/", func(req *gin.Context) {
		draftList := api.GetAll(db)
		fmt.Println(draftList)
		req.JSON(http.StatusOK, gin.H {
			"res": draftList.Value,
		})
	})
	router.Run(":3000")
}

