package main

import (
	"github.com/c0rrzin/router"
)

// DefRoutes is the function in which all routes are created
func DefRoutes() {
	router.DefRoute("GET", "/orders", GetOrdersHandler)
	router.DefRoute("GET", "/order", GetOrderHandler)
	router.DefRoute("POST", "/order", NewOrderHandler)
	router.DefRoute("POST", "/order/finish", FinishOrderHandler)
	router.DefRoute("POST", "/order/cancel", CancelOrderHandler)
	router.DefRoute("POST", "/order/approve", ApproveOrderHandler)
}
