package test

import (
	"log"
	"testing"
	"wanpei-backend/mapper"
	"wanpei-backend/models"
)

func TestName(t *testing.T) {
	conn := mapper.NewDbConn()
	var tags []models.UserGame
	conn.Conn.Where("ID = ?", 11).Find(&tags)
	log.Println("get tags = ", tags)
	allTags := map[string]uint{}
	log.Println(allTags["123"])
}
