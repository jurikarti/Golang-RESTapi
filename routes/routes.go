package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/your_project/controllers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/bookings", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/bookings/{id}", controllers.DeleteBooking).Methods("DELETE")
}