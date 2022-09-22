package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewApp(SessionMgr *cookie.Store) *gin.Engine {
	r := gin.Default()
	r.Use(sessions.Sessions("wanpei-session", *SessionMgr))
	r.Use(cors.Default())
	return r
}