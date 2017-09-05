package router

import (
	"grabber/news"
	"net/http"

	"github.com/gin-gonic/gin"
)

const GREETING_MESSAGE = `Welcome to the Grapper app!
  This app allows you to search for news in some area.
  Hit '/search/technology' and then hit '/result/technology' as an example.`

func mainHandler(c *gin.Context) {
	c.String(http.StatusOK, GREETING_MESSAGE)
}

func searchHandler(c *gin.Context) {
	categoryName := c.Param("category")
	news.SearchFor(categoryName)

	c.String(http.StatusOK, "Search for %s category is initialized.", categoryName)
}

func resultHandler(c *gin.Context) {
	categoryName := c.Param("category")
	topics := news.ResultFor(categoryName)

	c.JSON(http.StatusOK, topics)
}

func resetHandler(c *gin.Context) {
	news.ResetArchive()

	c.String(http.StatusOK, "The Archive is emptied successfully.")
}
