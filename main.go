/*
* Name: main.js
*	Description: Gin server and API endpoints configurations.
* Author: enderalansoy
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/*")
	api := router.Group("/")
	{
		api.GET("/", index)
		api.GET("/api", scrapePage)
	}
	router.Run(":8081")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func scrapePage(c *gin.Context) {
	url := c.Query("url")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	title, ok := scrape.Find(root, scrape.ByTag(atom.Title))
	if ok {
		c.JSON(http.StatusOK, gin.H{"title": scrape.Text(title)})
	}
}
