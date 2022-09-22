package services

import (
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

type Hub struct {
	HubMapper *mapper.Hub
}

func NewHub(hubMapper *mapper.Hub) *Hub {
	return &Hub{
		HubMapper: hubMapper,
	}
}

func (h *Hub) GetHub(name string) (*models.Hub, error) {
	hub, err := h.HubMapper.GetHub(name)
	if err != nil {
		return nil, err
	} else {
		return hub, nil
	}
}
