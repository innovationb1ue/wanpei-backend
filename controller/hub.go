package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/template"
	"wanpei-backend/models"
	"wanpei-backend/services"
	"wanpei-backend/utils"
)

type HubDeps struct {
	fx.In
	HubService *services.Hub
	Client     *services.Client
}

func HubRoutes(App *gin.Engine, Hub HubDeps) {
	App.GET("/hub/join", Hub.join)
	App.GET("/hub/users", Hub.users)
}

// join establish Websockets with users.
func (h *HubDeps) join(ctx *gin.Context) {
	// check require parameter: ID
	hubID := ctx.Query("ID")
	if hubID == "" {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "empty hub ID",
		})
		return
	}
	session := sessions.Default(ctx)
	userInterface := session.Get("user")
	if userInterface == nil {
		return
	}
	user := userInterface.(models.UserInsensitive)

	// get the hub
	hub, err := h.HubService.GetHub(hubID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "no such hub",
		})
		return
	}

	// check for user validation
	if ok := h.HubService.CheckUserValidity(hub.ID, &user); ok {
		log.Println("block one invalid user")
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "user is not allowed to connect to this hub",
		})
		return
	}

	// upgrade to Websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// todo: check headers for CSRF here.
			log.Println("Request Origin = ", r.Header.Get("Origin"))
			return true
		},
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// register client with HubService and start services
	h.Client.RegisterClient(hub, conn, &user)
}

func (h *HubDeps) users(ctx *gin.Context) {
	user, err := utils.ValidateLoginStatus(ctx)
	if err != nil {
		ctx.JSON(400, template.BaseError{
			Code:    -1,
			Message: "not login",
		})
		return
	}
	// get query param HubID
	HubID := ctx.Query("HubID")
	if HubID == "" {
		ctx.JSON(400, template.BaseError{
			Code:    -1,
			Message: "empty hub ID",
		})
		return
	}
	// get users in hub
	users := h.HubService.GetHubUsers(HubID)
	var userIDs []uint
	for _, u := range users {
		userIDs = append(userIDs, u.ID)
	}
	isValid := h.HubService.CheckValidUser(HubID, user)
	if !isValid {
		ctx.JSON(400, template.BaseError{
			Code:    -1,
			Message: "not authorized",
		})
		return
	}

	ctx.JSON(200, template.BaseResponse[[]models.UserSimple]{Code: 2, Data: users})

}
