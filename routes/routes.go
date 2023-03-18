package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jacewalker/tools/articles"
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

	if form.Name == "" || form.Email == "" || form.Message == "" {
		c.HTML(http.StatusBadRequest, "tmpl-contact.html", gin.H{
			"error": "Oops 😬, a field is missing. All except mobile number are required.",
		})
	} else {
		go contact.SendEmail(&form)
		c.HTML(http.StatusOK, "tmpl-contact.html", nil)
	}

}

func ArticlesRoute(c *gin.Context) {
	posts := articles.RetrieveArticles(c)

	c.HTML(http.StatusOK, "tmpl-articles.html", gin.H{
		"posts": posts,
	})
}

func ArticlePostRoute(c *gin.Context) {
	posts := articles.RetrieveArticles(c)
	var foundPost articles.Post
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		fmt.Println(err)
	}

	for _, post := range posts {
		if post.ID == postID {
			fmt.Println("Found Post! Post Title: ", post.Title)
			fmt.Println("Post ID: ", post.ID)
			foundPost = post
			break
		} else {
			fmt.Println("Could not find post for ", post.Title)
			// c.HTML(http.StatusNotFound, "tmpl-error404.html", gin.H{
			// 	"error": "Article not found",
			// })
		}
	}

	c.HTML(http.StatusOK, "tmpl-post.html", gin.H{
		"Title":   foundPost.Title,
		"Content": foundPost.Content,
		"ID":      foundPost.ID,
	})

}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "tmpl-error404.html", nil)
}
