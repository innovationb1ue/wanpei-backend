package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
)

func Run(App *gin.Engine) error {
	addr := "localhost:8096"
	serviceLn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		if err = App.RunListener(serviceLn); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf(`Server started on addr %s`, addr)

	return nil
}
