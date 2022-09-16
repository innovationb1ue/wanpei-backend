package models

import (
	"github.com/gin-contrib/sessions/cookie"
)

func NewSessionMgr() *cookie.Store {
	// todo: migrate secret to settings file
	CookieMgr := cookie.NewStore([]byte("todo: change secret here"))
	return &CookieMgr
}
