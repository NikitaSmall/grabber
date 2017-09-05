package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.New()

	r.GET("/", mainHandler)
	r.GET("/search/:category", searchHandler)
	r.GET("/result/:category", resultHandler)

	r.GET("/reset", resetHandler)

	return r
}
