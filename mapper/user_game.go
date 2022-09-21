package mapper

import (
	"gorm.io/gorm"
	"wanpei-backend/models"
)

type UserGame struct {
	DB *gorm.DB
}

func NewUserGame(db *DbConn) *UserGame {
	return &UserGame{
		DB: db.Conn,
	}
}

func (u *UserGame) GetUserGames(ID uint) []*models.UserGame {
	var tags []*models.UserGame
	u.DB.Where("ID = ?", ID).Find(&tags)
	return tags
}
