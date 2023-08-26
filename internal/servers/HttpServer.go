package servers

import (
	"context"
	"log"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	port    string
	context context.Context
	router  *mux.Router
}

func NewHttpServer(ctx context.Context) (*HttpServer, error) {
	return &HttpServer{
		port:    "8080",
		context: ctx,
	}, nil
}

func (server *HttpServer) Run(router func(s *HttpServer, r *mux.Router)) error {
	log.Println("HttpServer is running on port: ", server.port)

	router(server, server.router)
	return nil
}
