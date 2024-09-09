package controllers

import (
	"booking-api/models"
	"booking-api/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// создай
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := user.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	if err := repository.CreateUser(&user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// удали
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := repository.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := repository.DeleteUser(user); err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	if err := repository.DeleteBookingsByUserID(user.ID); err != nil {
		http.Error(w, "Error deleting user's bookings", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
