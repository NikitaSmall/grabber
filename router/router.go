package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.New()

	r.GET("/", mainHandler)

	return r
}
