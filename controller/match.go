package controller

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"go.uber.org/fx"
	"wanpei-backend/models"
)

// Match is the collection of dependencies
type Match struct {
	fx.In
	rdb *redis.Client
}

func MatchRoutes(App *gin.Engine, match Match) {
	App.POST("/match/start", match.Start)
}

func (m *Match) Start(ctx *gin.Context) {
	// get the current user

	session := sessions.Default(ctx)
	user := session.Get("user")
	if user == nil {
		ctx.JSON(404, gin.H{"code": -1, "message": "not logged in"})
		return
	}
	//todo: make a websocket conn with the client & handle the heartbeat thing
	userObj := user.(models.User)
	redisCtx := context.Background()
	m.rdb.LPush(redisCtx, "match:users", userObj.ID)
}
