package mapper

import "gorm.io/gorm"

type UserGame struct {
	Client *gorm.DB
}

func NewUserGame(db *DbConn) *Game {
	return &Game{
		DB: db.Conn,
	}
}

func (u *UserGame) GetUserGames(ID uint) []string {
	var tags []string
	u.Client.Where("ID = ?", ID).Find(&tags)
	return tags
}
