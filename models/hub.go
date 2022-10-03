package models

import (
	"context"
	"github.com/google/uuid"
	"log"
	"time"
)

// Hub maintains the set of active Client and broadcasts messages to the
// Client.
type Hub struct {
	// Registered Client.
	//todo: reformat this map to ID->Client
	Client map[*Client]bool

	// Inbound messages from the Client.
	Broadcast chan *ChatSocketMessage

	// ClientRegister requests from the Client.
	ClientRegister chan *Client

	// Unregister requests from Client.
	Unregister chan *Client

	Users map[uint]*UserInsensitive

	ID string

	AvailableUserID []uint
}

func NewHub() *Hub {
	return &Hub{
		Client:          make(map[*Client]bool),
		Broadcast:       make(chan *ChatSocketMessage),
		ClientRegister:  make(chan *Client),
		Unregister:      make(chan *Client),
		Users:           make(map[uint]*UserInsensitive),
		ID:              uuid.NewString(),
		AvailableUserID: []uint{},
	}
}

func (h *Hub) AppendAvailableUser(ID uint) {
	h.AvailableUserID = append(h.AvailableUserID, ID)
}

// Run is self-managed hub daemon thread. Will destroy if there is no one in the hub
func (h *Hub) Run(cancel context.CancelFunc) {
	// close empty hub
	ticker := time.NewTicker(100 * time.Second)
	for {
		select {
		case client := <-h.ClientRegister:
			log.Println("register one user & client", client)
			h.Client[client] = true
			h.Users[client.User.ID] = client.User
		case client := <-h.Unregister:
			if _, ok := h.Client[client]; ok {
				delete(h.Client, client)
				delete(h.Users, client.User.ID)
				close(client.Send)
			}
		// broadcast message
		case message := <-h.Broadcast:
			// iter through clients and send messages one by one
			for client := range h.Client {
				select {
				case client.Send <- message:
				// fail to send. close the client chan
				default:
					close(client.Send)
					delete(h.Client, client)
				}
			}
		// hub expired
		case <-ticker.C:
			{
				if len(h.Client) == 0 {
					log.Println("closed one hub")
					cancel()
					return
				}
			}
		}
	}
}
