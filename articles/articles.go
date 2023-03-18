package articles

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

type Post struct {
	ID      int
	Title   string
	Content template.HTML
}

// func ArticleNames() (postNames []string) {
// 	var posts []string

// 	files, err := os.ReadDir("./article-markdown/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		fmt.Println(file.Name())
// 		id := strings.Split(file.Name(), "#")
// 		fmt.Println("ID:", id[0])

// 		name := strings.Replace(id[1], ".md", "", 1)
// 		posts = append(posts, name)
// 	}
// 	return posts
// }

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

		idString := strings.Split(file.Name(), "#")
		id, err := strconv.Atoi(idString[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ID:", id)

		postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(fileByte)))
		postTitle := strings.Replace(idString[1], ".md", "", 1)
		posts = append(posts, Post{ID: id, Title: postTitle, Content: postHTML})
	}
	return posts
}
