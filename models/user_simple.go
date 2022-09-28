package models

type UserSimple struct {
	ID          uint   `json:"ID"`
	Nickname    string `json:"nickname"`
	AvatarURL   string `json:"avatar_url"`
	SteamCode   string `json:"steam_code"`
	Description string `json:"description"`
}
