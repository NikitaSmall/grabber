package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const GREETING_MESSAGE = `Welcome to the Grapper app!
  This app allows you to search for news in some area.
  Hit '/search/technology' and then hit '/result/technology' as an example.`

func mainHandler(c *gin.Context) {
	c.String(http.StatusOK, GREETING_MESSAGE)
}
