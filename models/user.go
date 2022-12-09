package models

import (
	"errors"
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

func (u User) ValidateChangeableUserFields() error {
	if !strings.Contains(u.Email, "@") {
		return errors.New("Email must have an symbol @. ")
	}
	if len(u.Email) < 6 {
		return errors.New("email must be longer than 6 chars")
	}
	if len(u.Email) > 20 {
		return errors.New("email too long")
	}
	if len(u.Nickname) > 20 {
		return errors.New("Nickname too long. ")
	}
	if len(u.SteamCode) > 30 {
		return errors.New("steam code too long. That's not possible")
	}
	return nil
}
