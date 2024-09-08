package routes

import (
	"booking-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	//User 
	r.HandleFunc("/users", controllers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{username}", controllers.DeleteUser).Methods(http.MethodDelete)

	//Booking 
	r.HandleFunc("/bookings", controllers.CreateBooking).Methods(http.MethodPost)
	r.HandleFunc("/bookings/{user_id}", controllers.GetBookings).Methods(http.MethodGet)

	return r
}
