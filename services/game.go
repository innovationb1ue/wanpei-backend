package services

import (
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

type Game struct {
	GameMapper *mapper.Game
}

func NewGame(userMapper *mapper.Game) *Game {
	return &Game{GameMapper: userMapper}
}

func (g *Game) GetAllGames() ([]*models.Game, error) {
	games, err := g.GameMapper.GetAllGames()
	if err != nil {
		return nil, err
	}
	return games, nil
}
