package services

import (
	"log"
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

func (h *Hub) CheckUserValidity(hubID string, user *models.UserInsensitive) bool {
	checks := []func(string, *models.UserInsensitive) bool{h.CheckDuplicateUser, h.CheckValidUser}
	for _, c := range checks {
		res := c(hubID, user)
		if !res {
			return false
		}
	}
	return true
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

func (h *Hub) CheckDuplicateUser(HubID string, user *models.UserInsensitive) bool {
	users := h.GetHubUsers(HubID)
	for _, u := range users {
		if u.ID == user.ID {
			return true
		}
	}
	return false
}

func (h *Hub) CheckValidUser(hubID string, user *models.UserInsensitive) bool {
	hub, err := h.GetHub(hubID)
	if err != nil {
		log.Println("hub not exist")
	}
	availableUsers := hub.AvailableUserID
	for _, u := range availableUsers {
		if u == user.ID {
			return true
		}
	}
	return false
}
