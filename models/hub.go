package models

import (
	"github.com/google/uuid"
	"log"
	"time"
)

// Hub maintains the set of active Client and broadcasts messages to the
// Client.
type Hub struct {
	// Registered Client.
	Client map[*Client]bool

	// Inbound messages from the Client.
	Broadcast chan []byte

	// Register requests from the Client.
	Register chan *Client

	// Unregister requests from Client.
	Unregister chan *Client

	ID string
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Client:     make(map[*Client]bool),
		ID:         uuid.NewString(),
	}
}

func (h *Hub) Run() {
	// close empty hub
	ticker := time.NewTicker(100 * time.Second)
	for {
		select {
		case client := <-h.Register:
			h.Client[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Client[client]; ok {
				delete(h.Client, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Client {
				select {
				case client.Send <- message:
				// fail to send. close the client chan
				default:
					close(client.Send)
					delete(h.Client, client)
				}
			}
		case <-ticker.C:
			{
				log.Println("closed one hub")
				if len(h.Client) == 0 {
					return
				}
			}
		}

	}
}
