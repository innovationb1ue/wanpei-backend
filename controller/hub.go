package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/template"
	"wanpei-backend/models"
	"wanpei-backend/services"
)

type Room struct {
	fx.In
	Hub *services.Hub
}

func HubRoutes(App *gin.Engine, Room Room) {
	App.GET("/hub", Room.ConnectHub)
}

// ConnectHub establish Websockets with users.
func (r *Room) ConnectHub(ctx *gin.Context) {
	// check require parameter: ID
	hubID := ctx.Query("ID")
	if hubID == "" {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "empty ConnectHub ID",
		})
		return
	}
	// get the ConnectHub
	hub, err := r.Hub.GetHub(hubID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, template.BaseError{
			Code:    -1,
			Message: "no such ConnectHub",
		})
		return
	}

	// upgrade to Websocket
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// todo: check headers for CSRF here.
			return true
		},
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// spawn server-side client object
	client := &models.Client{Hub: hub, Conn: conn, Send: make(chan *models.ChatSocketMessage, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
