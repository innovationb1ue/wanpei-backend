package server

import "github.com/gin-gonic/gin"

func NewApp() *gin.Engine {
	// todo: add custom middlewares
	r := gin.Default()
	// todo: divide handlers into different files based on functions
	r.POST("/user/create", func(context *gin.Context) {

	})
	return r
}
