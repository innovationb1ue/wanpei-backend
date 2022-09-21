package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"wanpei-backend/server"
)

func NewSessionStore(settings *server.Settings) *cookie.Store {
	options := sessions.Options{
		Path:     "",
		Domain:   "",
		MaxAge:   1000,
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	}
	CookieMgr := cookie.NewStore([]byte(settings.Secret))
	CookieMgr.Options(options)
	return &CookieMgr
}
