package models

import (
	"gorm.io/gorm"
	"strings"
)

type UserInsensitive struct {
	gorm.Model  `json:"Gorm.Model"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Games       string `json:"-"`
	SteamCode   string `json:"steam_code"`
	Description string `json:"description"`
}

func (UserInsensitive) TableName() string {
	return "users"
}
func (u UserInsensitive) ValidateChangeableUserFields() bool {
	if !strings.Contains(u.Email, "@") || len(u.Email) < 8 || len(u.Email) > 20 {
		return false
	}
	if len(u.Nickname) > 20 {
		return false
	}
	if len(u.SteamCode) > 30 {
		return false
	}
	return true
}
