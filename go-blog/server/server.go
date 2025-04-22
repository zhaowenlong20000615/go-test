package server

import (
	"go-test/go-blog/routers"
	"log"
	"net/http"
)

var App = &AppServer{}

type AppServer struct {
}

func (app *AppServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	routers.Routes()
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Error starting server: %s\n", err)
	}
}
