package main

import (
	"booking-api/config"
	"booking-api/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDatabase()

	r := routes.RegisterRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
