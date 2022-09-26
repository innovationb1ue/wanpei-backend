package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"strings"
	"wanpei-backend/controller/template"
	"wanpei-backend/repo"
	"wanpei-backend/services"
	"wanpei-backend/utils"
)

// Match is the collection of dependencies
type Match struct {
	fx.In
	MatchService  *services.Match
	TokenService  *services.Token
	SocketService *services.Socket
	UserGameRepo  *repo.QueueUserGame
}

func MatchRoutes(App *gin.Engine, match Match) {
	App.POST("/match/start", match.Start)
	App.POST("/match/stop", match.Stop)
	App.GET("/match/socket", match.Socket)
	App.GET("/match/current", match.Current)
}

func (m *Match) Start(ctx *gin.Context) {
	// validate login status
	user, err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		ctx.JSON(404, template.BaseResponse[any]{
			Code:    -1,
			Message: "Not logged in",
			Data:    nil,
		})
		return
	}
	jsonHolder := struct {
		SelectedGame []int `json:"selectedGame"`
	}{}
	err = ctx.ShouldBindJSON(&jsonHolder)
	if err != nil {
		return
	}
	token := m.TokenService.GenerateRandom(user.ID)
	ctx.JSON(200, template.BaseResponse[string]{
		Code:    0,
		Message: "ok",
		Data:    token,
	})
}

func (m *Match) Stop(ctx *gin.Context) {
	user, err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		ctx.JSON(404, template.BaseResponse[any]{
			Code:    -1,
			Message: "Not logged in",
			Data:    nil,
		})
		return
	}
	m.MatchService.RemoveFromQueue(user.ID)
	ctx.JSON(200, template.BaseResponse[any]{
		Code:    1,
		Message: "ok",
		Data:    nil,
	})
	log.Println("one stopped match making, now queue = ", m.MatchService.RedisMapper.GetAllFromQueue())
}

func (m *Match) Socket(ctx *gin.Context) {
	// check valid token
	token := ctx.Query("auth")
	userID, err := m.TokenService.GetUserID(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, template.BaseErrorResponse())
		return
	}
	// get selected games
	GamesStr, isSelected := ctx.GetQuery("selectedGame")
	if !isSelected {
		ctx.JSON(http.StatusBadRequest, template.BaseErrorResponse())
		return
	}
	gameIdStrArr := strings.Split(GamesStr, ",")
	GamesInt := utils.AllAsInt[int](gameIdStrArr)
	m.UserGameRepo.UserGame[userID] = GamesInt

	// upgrade to websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// todo: check headers for CSRF here.
			return true
		},
	}
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "Upgrade to socket failed",
		})
		return
	}

	// add socket to collection & put user in queue
	m.SocketService.AppendSocket(userID, ws)

	// debug message
	log.Println("Now queue = ", m.MatchService.RedisMapper.GetAllFromQueue())

}

func (m *Match) Current(ctx *gin.Context) {
	user := sessions.Default(ctx).Get("user")
	if user != nil {
		ctx.JSON(200, gin.H{"data": user})
	} else {
		ctx.JSON(200, "failed")
	}
}
