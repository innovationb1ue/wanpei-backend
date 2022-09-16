package controller

import "github.com/gin-gonic/gin"

func ping(App *gin.Engine) {
	App.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong!")
	})
}
