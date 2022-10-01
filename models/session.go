package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"wanpei-backend/server"
)

func NewSessionStore(settings *server.Settings) *cookie.Store {
	CookieMgr := cookie.NewStore([]byte(settings.CookieSecret))
	CookieMgr.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   60 * 60,
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})
	return &CookieMgr
}
