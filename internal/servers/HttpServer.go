package servers

import (
	"context"
	"log"
)

type HttpServer struct {
	port    string
	context context.Context
}

func NewHttpServer(ctx context.Context) (*HttpServer, error) {
	return &HttpServer{
		port:    "8080",
		context: ctx,
	}, nil
}

func (server *HttpServer) Run() error {
	log.Println("HttpServer is running on port: ", server.port)
	return nil
}
