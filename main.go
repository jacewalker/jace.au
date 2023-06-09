package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jacewalker/tools/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.LoadHTMLGlob("./views/*/*.html")
	r.StaticFile("/assets/imgs/jace-cartoon.jpeg", "./assets/imgs/jace-cartoon.jpeg")
	r.StaticFile("/assets/imgs/jace-gage.jpeg", "./assets/imgs/jace-gage.jpeg")
	r.StaticFile("/assets/imgs/jace-nat.jpeg", "./assets/imgs/jace-nat.jpeg")
	r.StaticFile("/assets/imgs/jace-plane.jpeg", "./assets/imgs/jace-plane.jpeg")
	r.StaticFile("/assets/imgs/jace-tesla.jpeg", "./assets/imgs/jace-tesla.jpeg")
	r.StaticFile("/assets/output.css", "./assets/output.css")
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
