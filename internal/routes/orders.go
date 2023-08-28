package routes

import (
	"github.com/gorilla/mux"
	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/handlers/orders"
	"github.com/gt2rz/change-status_order_demo/delivery_api/internal/servers"
)

func OrderRoutes(server *servers.HttpServer, router *mux.Router) {
	router.HandleFunc("/api/orders/v1/change_status/{orderId}", orders.ChangeStatusHandler(server)).Methods("PATCH")
}
