package controllers

import (
	"booking-api/models"
	"booking-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking with the input payload
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param booking body models.Booking true "Create Booking"
// @Success 201 {object} models.Booking
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Error creating booking"
// @Router /bookings [post]
func CreateBooking(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := booking.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateBooking(&booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// GetBookings godoc
// @Summary Get bookings by user ID
// @Description Retrieve bookings for a specific user by user ID
// @Tags bookings
// @Produce  json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Booking
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Error retrieving bookings"
// @Router /bookings/{user_id} [get]
func GetBookings(c *gin.Context) {
	userIDStr := c.Param("user_id")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	bookings, err := repository.GetBookingsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
