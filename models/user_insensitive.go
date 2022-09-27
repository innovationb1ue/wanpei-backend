package models

import "gorm.io/gorm"

type UserInsensitive struct {
	gorm.Model `json:"Gorm.Model"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	Games      string `json:"games"`
	SteamCode  string `json:"steam_code"`
}

func (UserInsensitive) TableName() string {
	return "users"
}
