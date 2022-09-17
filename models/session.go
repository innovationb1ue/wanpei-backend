package models

import (
	"github.com/gin-contrib/sessions/cookie"
	"wanpei-backend/server"
)

func NewSessionStore(settings *server.Settings) *cookie.Store {
	// todo: migrate secret to settings file
	CookieMgr := cookie.NewStore([]byte(settings.Secret))
	return &CookieMgr
}
