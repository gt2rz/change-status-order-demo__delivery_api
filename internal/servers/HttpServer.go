package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
		router:  mux.NewRouter(),
	}, nil
}

func (server *HttpServer) Run(router func(s *HttpServer, r *mux.Router)) error {
	log.Println("HttpServer is running on port: ", server.port)

	router(server, server.router)

	err := http.ListenAndServe(fmt.Sprintf(":%s", server.port), server.router)
	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}
