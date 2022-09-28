package models

import (
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model  `json:"Gorm.Model"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Nickname    string `json:"nickname"`
	Games       string `json:"games"`
	SteamCode   string `json:"steam_code"`
	UserRole    int    `json:"user_role"`
	Description string `json:"description"`
}

// TableName indicates the target table for GORM
func (u User) TableName() string {
	return "users"
}

func (u User) ValidateChangeableUserFields() bool {
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
