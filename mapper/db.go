package mapper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DbConn struct {
	Conn *gorm.DB
}

func NewDbConn() *DbConn {
	dsn := "b1ue:TheHardestOne@tcp(host.docker.internal:3306)/wanpei?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection failed. ")
	}
	return &DbConn{Conn: db}
}
