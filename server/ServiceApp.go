package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewApp(SessionMgr *cookie.Store) *gin.Engine {
	// todo: add custom middlewares or options
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", *SessionMgr))
	return r
}
