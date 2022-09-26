package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/template"
	"wanpei-backend/services"
)

type HubDeps struct {
	fx.In
	Hub    *services.Hub
	Client *services.Client
}

func HubRoutes(App *gin.Engine, Hub HubDeps) {
	App.GET("/hub", Hub.JoinHub)
}

// JoinHub establish Websockets with users.
func (r *HubDeps) JoinHub(ctx *gin.Context) {
	// check require parameter: ID
	hubID := ctx.Query("ID")
	if hubID == "" {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "empty hub ID",
		})
		return
	}
	// get the JoinHub
	hub, err := r.Hub.GetHub(hubID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "no such JoinHub",
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
	// register client with Hub and start services
	r.Client.RegisterClient(hub, conn)
}
