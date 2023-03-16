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

	r := gin.Default()
	r.LoadHTMLGlob("views/*/*.html")
	r.StaticFile("/assets/imgs/jace-cartoon.jpeg", "./assets/imgs/jace-cartoon.jpeg")
	r.GET("/", routes.IndexRoute)
	r.GET("/about", routes.AboutMeRoute)
	r.GET("/projects", routes.ProjectsRoute)
	r.GET("/contact", routes.ContactRoute)
	r.POST("/contact", routes.ContactSubmissionRoute)
	r.NoRoute(routes.NotFound)
	r.Run(":8080")
}
