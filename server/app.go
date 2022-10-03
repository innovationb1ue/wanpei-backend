package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewApp(SessionMgr *cookie.Store) *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("wanpei-session", *SessionMgr))
	r.Use(cors.Default())
	r.Use(gin.CustomRecovery(Recovery))
	r.Use(gin.Logger())

	return r
}
