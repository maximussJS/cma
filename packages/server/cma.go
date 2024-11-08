package server

import (
	"cma/packages/config"
	"cma/packages/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CMAServer struct {
	Router *mux.Router
}

func NewCMAServer() *CMAServer {
	server := &CMAServer{
		Router: mux.NewRouter(),
	}

	server.initRoutes()

	return server
}

func (s *CMAServer) initRoutes() {
	internalHandler := handlers.NewInternalHandler()
	webhookHandler := handlers.NewWebhookHandler()

	internalHandler.RegisterRoutes(s.Router)
	webhookHandler.RegisterRoutes(s.Router)
}

func (s *CMAServer) Start() {
	log.Printf("Starting server on port %s", config.GlobalConfig.Port)
	if err := http.ListenAndServe(config.GlobalConfig.Port, s.Router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
