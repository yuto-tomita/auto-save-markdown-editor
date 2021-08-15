package main

import (
	"net/http"

	"myapp/pkg/api"
	"myapp/pkg/db"
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
		var json model.Draft
		
		if err := req.ShouldBindJSON(&json); err != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		api.SaveMarkdown(db, json)

		req.JSON(http.StatusOK, gin.H{json.Title: json.Markdown_text})
	})

	router.GET("/", func(req *gin.Context) {
		draftList := api.GetAll(db)
		req.JSON(http.StatusOK, gin.H {
			"res": draftList.Value,
		})
	})

	router.DELETE("/draft/:id", func(req *gin.Context) {
		api.DeleteDraft(db, req.Param("id"))
	})
	router.Run(":3000")
}

