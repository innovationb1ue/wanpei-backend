package test

import (
	"testing"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

func TestGetAllGames(t *testing.T) {
	conn := models.NewDbConn()
	gameObj := mapper.Game{DB: conn.Conn}

	games, err := gameObj.GetAllGames()
	if err != nil {
		t.Fatal("failed to get all games", err)
	}
	var gameNames []string
	for _, i := range games {
		gameNames = append(gameNames, i.GameName)
	}

	t.Logf("Get all games = %s", gameNames)
}
