package main

import (
	"log"

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
	r.LoadHTMLGlob("./views/*/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", routes.IndexRoute)
	r.GET("/about", routes.AboutMeRoute)
	r.GET("/projects", routes.ProjectsRoute)
	r.GET("/posts", routes.ArticlesRoute)
	r.GET("/posts/:postID", routes.ArticlePostRoute)
	r.GET("/contact", routes.ContactRoute)
	r.POST("/contact", routes.ContactSubmissionRoute)
	r.NoRoute(routes.NotFound)
	r.Run(":8080")
}
