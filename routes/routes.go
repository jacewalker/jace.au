package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacewalker/tools/contact"
)

func IndexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "tmpl-home.html", gin.H{
		"title": "Jace's Tools",
	})
}

func AboutMeRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "tmpl-aboutme.html", gin.H{
		"title": "Jace's Tools",
	})
}

func ContactRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "tmpl-contact.html", gin.H{
		"title": "Jace's Tools",
	})
}

func ProjectsRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "tmpl-projects.html", gin.H{
		"title": "Jace's Tools",
	})
}

func ContactSubmissionRoute(c *gin.Context) {
	var form contact.Form
	form.Name = c.PostForm("name")
	form.Email = c.PostForm("email")
	form.Phone = c.PostForm("phone")
	form.Message = c.PostForm("message")

	go contact.SendEmail(&form)
	c.HTML(http.StatusOK, "tmpl-contact.html", nil)
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "tmpl-error404.html", nil)
}
