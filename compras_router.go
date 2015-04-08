package main

import (
	"fmt"
	"net/http"

	"github.com/c0rrzin/router"
)

// DefRoutes is the function in which all routes are created
func DefRoutes() {
	router.DefRoute("GET", "/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List Orders")
	})
	router.DefRoute("GET", "/order", GetOrderHandler)
	router.DefRoute("POST", "/order", NewOrderHandler)
	router.DefRoute("PUT", "/order/finish", FinishOrderHandler)
	router.DefRoute("PUT", "/order/cancel", CancelOrderHandler)
	router.DefRoute("PUT", "/order/approve", ApproveOrderHandler)
}
