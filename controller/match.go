package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"strings"
	"wanpei-backend/controller/template"
	"wanpei-backend/models"
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
	matchGroup := App.Group("/match")
	matchGroup.Use(ValidateLoginStatus)
	matchGroup.POST("/start", match.Start)
	matchGroup.POST("/stop", match.Stop)
	matchGroup.GET("/socket", match.Socket)
	matchGroup.GET("/current", match.Current)
	matchGroup.GET("/count", match.PlayerCount)
}

func (m *Match) Start(ctx *gin.Context) {
	userInsensitive := ctx.MustGet("user").(models.UserInsensitive)
	// create an empty struct to receive query params
	jsonHolder := struct {
		SelectedGame []int `json:"selectedGame"`
	}{}
	err := ctx.ShouldBindJSON(&jsonHolder)
	if err != nil {
		return
	}
	token := m.TokenService.GenerateRandom(userInsensitive.ID)
	ctx.JSON(200, template.BaseResponse[string]{
		Code:    0,
		Message: "ok",
		Data:    token,
	})
}

func (m *Match) Stop(ctx *gin.Context) {
	userInsensitive := ctx.MustGet("user").(models.UserInsensitive)
	m.MatchService.RemoveFromQueue(userInsensitive.ID)
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
	userObj, ok := ctx.Get("user")
	log.Print("context user = ", userObj)
	if ok && userObj != nil {
		user := userObj.(models.UserInsensitive)
		ctx.JSON(200, gin.H{"data": user})
	} else {
		return
	}
}

func (m *Match) PlayerCount(ctx *gin.Context) {
	queueLength, err := m.MatchService.QueueLength()
	if err != nil {
		ctx.JSON(200, gin.H{"count": 0})
		return
	}
	ctx.JSON(200, gin.H{"count": queueLength})
}
