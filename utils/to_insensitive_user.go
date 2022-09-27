package utils

import (
	"gorm.io/gorm"
	"wanpei-backend/models"
)

func ToInsensitiveUser(user *models.User) *models.UserInsensitive {
	return &models.UserInsensitive{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Username:  user.Username,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Games:     user.Games,
		SteamCode: user.SteamCode,
	}
}
