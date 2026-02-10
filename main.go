package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jacewalker/tools/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.LoadHTMLGlob("./views/*/*.html")

	// Static assets with long cache headers (1 year) and immutable for fingerprinted files
	r.Group("/assets").Use(func(c *gin.Context) {
		ext := strings.ToLower(filepath.Ext(c.Request.URL.Path))
		switch ext {
		case ".css", ".js", ".woff2", ".woff", ".ttf":
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		case ".jpg", ".jpeg", ".png", ".webp", ".gif", ".svg", ".ico":
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		default:
			c.Header("Cache-Control", "public, max-age=86400")
		}
		c.Next()
	}).Static("/", "./assets")

	r.GET("/", routes.IndexRoute)
	r.GET("/services", routes.ServicesRoute)
	r.GET("/about", routes.AboutMeRoute)
	r.GET("/projects", routes.ProjectsRoute)
	r.GET("/posts", routes.ArticlesRoute)
	r.GET("/posts/:postID", routes.ArticlePostRoute)
	r.GET("/contact", routes.ContactRoute)
	r.POST("/contact", routes.ContactSubmissionRoute)
	r.GET("/tools/qr-generator", routes.QRGeneratorRoute)
	r.GET("/tools/salary-sacrifice", routes.SalarySacrificeRoute)
	r.NoRoute(routes.NotFound)
	r.Run(":8080")
}
