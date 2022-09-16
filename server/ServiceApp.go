package server

import "github.com/gin-gonic/gin"

func NewApp() *gin.Engine {
	// todo: add custom middlewares or options
	r := gin.Default()
	return r
}
