package controllers

import (
	"booking-api/models"
	"booking-api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// создай
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := booking.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.CreateBooking(&booking); err != nil {
		http.Error(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
 
// дай
func GetBookings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	bookings, err := repository.GetBookingsByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Error retrieving bookings", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookings)
}
