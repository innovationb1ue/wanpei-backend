package mapper

import (
	"errors"
	"wanpei-backend/models"
	"wanpei-backend/repo"
)

type Hub struct {
	HubRepo *repo.Hub
}

func NewHub(hubRepo *repo.Hub) *Hub {
	return &Hub{HubRepo: hubRepo}
}

func (h *Hub) RegisterNewHub(hub *models.Hub) {
	h.HubRepo.Hubs[hub.ID] = hub
}

func (h *Hub) DeleteHub(name string) {
	delete(h.HubRepo.Hubs, name)
}

func (h *Hub) GetHub(name string) (*models.Hub, error) {
	if hub := h.HubRepo.Hubs[name]; hub == nil {
		return nil, errors.New("no such hub")
	} else {
		return hub, nil
	}
}

func (h *Hub) GetHubValidUsers(name string) []uint {
	hub, err := h.GetHub(name)
	if err != nil {
		return nil
	}
	return hub.AvailableUserID
}
