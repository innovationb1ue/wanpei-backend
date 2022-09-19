package controller

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"wanpei-backend/controller/types"
	"wanpei-backend/services"
)

// User include all the services that would be used when handling the req
type Game struct {
	fx.In
	GameServices *services.Game
	SessionMgr   *cookie.Store
}

func GameRoutes(App *gin.Engine, game Game) {
	App.GET("/game/all", game.GetAll)
	App.POST("/game/all", game.GetAll)
}

func (g *Game) GetAll(ctx *gin.Context) {
	games, err := g.GameServices.GetAllGames()
	if err != nil {
		ctx.JSON(400, gin.H{"meesage": "fail to get all games"})
		return
	}
	ctx.JSON(200, types.BaseResponse{
		Code:    200,
		Message: "ok",
		Data:    games,
	})

}
