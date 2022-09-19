package controller

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"net/http"
	"wanpei-backend/controller/types"
	"wanpei-backend/models"
	"wanpei-backend/utils"
)

// Match is the collection of dependencies
type Match struct {
	fx.In
	rdb *redis.Client
}

func MatchRoutes(App *gin.Engine, match Match) {
	App.POST("/match/start", match.Start)
	App.POST("/match/socket")
}

func (m *Match) Start(ctx *gin.Context) {
	// validate login status
	err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		return
	}
	// get the user
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

// Socket establishes a websocket with the client
func (m *Match) Socket(ctx *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	w, r := ctx.Writer, ctx.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseErrorResponse())
	}
	// placeholder
	conn.Close()
	//todo: make a websocket collection and append websockets to it
}
