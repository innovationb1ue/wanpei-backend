package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/types"
	"wanpei-backend/services"
	"wanpei-backend/utils"
)

// Match is the collection of dependencies
type Match struct {
	fx.In
	MatchService  *services.Match
	TokenService  *services.Token
	SocketService *services.Socket
}

func MatchRoutes(App *gin.Engine, match Match) {
	App.POST("/match/start", match.Start)
	App.GET("/match/socket", match.Socket)
}

func (m *Match) Start(ctx *gin.Context) {
	// validate login status, will
	user, err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		ctx.JSON(404, types.BaseResponse[any]{
			Code:    -1,
			Message: "Not logged in",
			Data:    nil,
		})
		return
	}
	token := m.TokenService.GenerateRandom(user.ID)
	ctx.JSON(200, types.BaseResponse[string]{
		Code:    0,
		Message: "ok",
		Data:    token,
	})
}

func (m *Match) Socket(ctx *gin.Context) {
	// check valid token
	token := ctx.Query("auth")
	userID, err := m.TokenService.GetUserID(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseErrorResponse())
		return
	}
	// upgrade to websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// todo: check headers for CSRF here.
			return true
		},
	}
	w, r := ctx.Writer, ctx.Request
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.BaseError{
			Code:    -1,
			Message: "Upgrade to socket failed",
		})
		return
	}
	// add socket to collection
	m.SocketService.AppendSocket(userID, ws)
	// write a success message back to client
	if err = ws.WriteJSON(gin.H{"message": "ok from server"}); err != nil {
		return
	}
	// start heartbeat
	go m.MatchService.StartHeartbeat(userID)
	// append user to queue
	_, err = m.MatchService.AppendToQueue(userID)
	if err != nil {
		log.Println("Append user to Redis queue failed.")
		return
	}

	log.Println("Now queue = ", m.MatchService.RedisMapper.GetAllFromQueue())

}
