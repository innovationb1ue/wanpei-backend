package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Run(App *gin.Engine, settings *Settings) error {
	addr := settings.Addr
	go func() {
		if err := App.RunTLS(addr, "./cert/scs1664441599798_wanpei.top_server.crt", "./cert/scs1664441599798_wanpei.top_server.key"); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf(`Server started on addr %s`, addr)

	return nil
}
