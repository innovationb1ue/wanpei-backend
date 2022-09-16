package test

import (
	"fmt"
	"log"
	"testing"
	"wanpei-backend/models"
)

func TestAddUser(t *testing.T) {
	user := models.User{
		Username: "b1ue",
		Password: "123123123",
		Nickname: "b1ue nickname",
		Email:    "123@456.com",
		Games:    fmt.Sprintf("%v", []string{"CSGO", "DOTA2"}),
	}
	conn := models.NewDbConn()
	conn.Conn.Create(&user)
	log.Print("Created user, id = ", user.ID)
}
