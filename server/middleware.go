package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// global middlewares
// **********************

func Recovery(ctx *gin.Context, recovered any) {
	if err, ok := recovered.(string); ok {
		log.Println("Recovered from err :", err)
	}
	ctx.AbortWithStatus(http.StatusInternalServerError)
}
