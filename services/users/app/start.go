package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mostafababaii/go-micro/services/users/app/routers"
	"github.com/mostafababaii/go-micro/services/users/config"
)

func StartApplication() {
	address := fmt.Sprintf(
		"%s:%d",
		config.ServerHost,
		config.ServerPort,
	)

	server := http.Server{
		Addr:    address,
		Handler: routers.Default,
	}

	var c chan bool

	go func(s *http.Server) {
		s.ListenAndServe()
		c <- true
	}(&server)

	log.Println(fmt.Sprintf("Started server on %s", address))

	<-c
}
