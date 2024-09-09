package repository

import (
	"booking-api/models"
	"booking-api/config"
	"errors"
)

func CreateBooking(booking *models.Booking) error {
	// Валидация данных бронирования
	if err := booking.Validate(); err != nil {
		return err
	}

	if booking.StartTime.After(booking.EndTime) {
		return errors.New("start time must be before end time")
	}

	return config.DB.Create(booking).Error
}

func GetBookingsByUserID(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	err := config.DB.Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func DeleteBookingsByUserID(userID uint) error {
	return config.DB.Where("user_id = ?", userID).Delete(&models.Booking{}).Error
}
