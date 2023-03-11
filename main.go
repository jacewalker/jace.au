package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*/*.html")
	router.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tmpl-article.html", gin.H{
			"title": "Jace's Tools",
		})
	})
	router.GET("/article2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tmpl-form.html", gin.H{
			"title": "Jace's Tools",
		})
	})
	router.Run(":8080")
}
