package models

type Game struct {
	ID              int    `json:"ID"`
	GameName        string `json:"game_name"`
	GameDescription string `json:"game_description"`
}

// TableName indicates the target table for GORM
func (g Game) TableName() string {
	return "games"
}
