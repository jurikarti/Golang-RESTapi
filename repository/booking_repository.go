package repository

import (
	"booking-api/models"
	"booking-api/config"
)

func CreateBooking(booking *models.Booking) error {
	return config.DB.Create(booking).Error
}

func GetBookingsByUserID(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func DeleteBookingsByUserID(userID uint) error {
	return config.DB.Where("user_id = ?", userID).Delete(&models.Booking{}).Error
}