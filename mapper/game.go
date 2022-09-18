package mapper

import (
	"gorm.io/gorm"
	"log"
	"wanpei-backend/models"
)

type Game struct {
	DB *gorm.DB
}

func NewGame(db *models.DbConn) *Game {
	return &Game{
		DB: db.Conn,
	}
}

func (g *Game) AddGame(game *models.Game) error {
	res := g.DB.Create(&game)
	if res.Error != nil {
		log.Fatal("Add game failed")
		return res.Error
	}
	return nil
}

func (g *Game) GetAllGames() ([]*models.Game, error) {
	var Games []*models.Game
	res := g.DB.Find(&Games)
	if res.Error != nil {
		log.Fatal("Get all games failed")
		return nil, res.Error
	}
	return Games, nil
}
