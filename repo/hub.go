package repo

import "wanpei-backend/models"

type Hub struct {
	Hubs map[string]*models.Hub // this holds all the active hubs
}

func NewHub() *Hub {
	return &Hub{Hubs: make(map[string]*models.Hub)}
}
