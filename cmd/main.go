package main

import (
	"context"
	"log"

	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/servers"
)

func init(){
	log.Println("Initializing...!")
}

func main() {
	httpServer, err := servers.NewHttpServer(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if err := httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
