package articles

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type Post struct {
	Title   string
	Content template.HTML
}

func ArticleNames() (postNames []string) {
	var posts []string

	files, err := os.ReadDir("./article-markdown/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		name := strings.Replace(file.Name(), ".md", "", 1)
		posts = append(posts, name)
	}
	return posts
}

func RetrieveArticles(c *gin.Context) []Post {
	var posts []Post

	dir, err := os.ReadDir("./article-markdown/")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range dir {
		fileByte, err := os.ReadFile("./article-markdown/" + file.Name())
		if err != nil {
			fmt.Println("ERROR: ", err)
			c.HTML(http.StatusNotFound, "tmpl-error404.html", nil)
			c.Abort()
			return []Post{}
		}
		postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(fileByte)))
		postTitle := strings.Replace(file.Name(), ".md", "", 1)
		posts = append(posts, Post{Title: postTitle, Content: postHTML})
	}
	return posts
}
