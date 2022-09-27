package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"Gorm.Model"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Nickname   string `json:"nickname"`
	Games      string `json:"games"`
	SteamCode  string `json:"steam_code"`
	UserRole   int    `json:"user_role"`
}

// TableName indicates the target table for GORM
func (g User) TableName() string {
	return "users"
}
