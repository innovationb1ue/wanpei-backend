package repo

import "wanpei-backend/models"

type Hub struct {
	Hubs map[string]*models.Hub
}

func NewHub() *Hub {
	return &Hub{Hubs: map[string]*models.Hub{}}
}
