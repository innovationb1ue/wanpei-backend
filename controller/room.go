package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
	"wanpei-backend/controller/types"
	"wanpei-backend/repo"
)

type Room struct {
	fx.In
	Hub *repo.Hub
}

func RoomRoutes(App *gin.Engine, Room *Room) {
	App.GET("/room", Room.room)
}

// room establish Websockets with users and call services to start.
func (r *Room) room(ctx *gin.Context) {
	// retrieve the Hub ID
	RoomID := ctx.Query("ID")
	if RoomID == "" {
		ctx.JSON(http.StatusBadRequest, types.BaseError{
			Code:    -1,
			Message: "empty room ID",
		})
	}
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// todo: Create client obj and register to Hub

	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

}
