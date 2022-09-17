package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"Gorm.Model"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Nickname   string `json:"nickname" `
	Games      string `json:"games"`
	UserRole   string `json:"user-role"`
}
