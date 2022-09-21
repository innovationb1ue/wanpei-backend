package mapper

import "gorm.io/gorm"

type UserGame struct {
	DB *gorm.DB
}

func NewUserGame(db *DbConn) *UserGame {
	return &UserGame{
		DB: db.Conn,
	}
}

func (u *UserGame) GetUserGames(ID uint) []string {
	var tags []string
	u.DB.Where("ID = ?", ID).Find(&tags)
	return tags
}
