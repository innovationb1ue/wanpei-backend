package models

type UserGame struct {
	ID     uint `json:"ID,omitempty"` // this ID is user ID
	GameID int  `json:"game_id" `     // Game ID
}
