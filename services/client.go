package services

import (
	"github.com/gorilla/websocket"
	"wanpei-backend/models"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) RegisterClient(hub *models.Hub, conn *websocket.Conn, user *models.UserInsensitive) {
	// spawn server-side client object
	client := &models.Client{Hub: hub, Conn: conn, Send: make(chan *models.ChatSocketMessage, 256), User: user}
	client.Hub.ClientRegister <- client
	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
