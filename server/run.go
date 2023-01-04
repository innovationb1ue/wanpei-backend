package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Run(App *gin.Engine, settings *Settings, env *Env) error {
	addr := settings.Addr
	isHttps := settings.https
	if isHttps {
		go func() {
			log.Println(env.Vars["cert_file"], env.Vars["cert_key"])
			if err := App.RunTLS(addr, env.Vars["cert_file"], env.Vars["cert_key"]); err != nil {
				log.Fatal(err)
			}
		}()
	} else {
		go func() {
			if err := App.Run(addr); err != nil {
				log.Fatal(err)
			}
		}()
	}

	log.Printf(`Server started on addr %s`, addr)

	return nil
}
