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

func (h *Hub) GetHub(ID string) (*models.Hub, error) {
	hub, err := h.HubMapper.GetHub(ID)
	if err != nil {
		return nil, err
	} else {
		return hub, nil
	}
}

func (h *Hub) GetHubUsers(ID string) []models.UserSimple {
	hub, err := h.HubMapper.GetHub(ID)
	if err != nil {
		return nil
	}
	var users []models.UserSimple
	for _, u := range hub.Users {
		users = append(users, models.UserSimple{
			ID:          u.ID,
			Nickname:    u.Nickname,
			AvatarURL:   "",
			SteamCode:   u.SteamCode,
			Description: u.Description,
		})
	}
	return users

}
