package models

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "time"
)

type Booking struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    UserID    uint      `json:"user_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
}

// Validate проверяет корректность полей модели Booking
func (booking *Booking) Validate() error {
    return validation.ValidateStruct(booking,
        validation.Field(&booking.StartTime, validation.Required),
        validation.Field(&booking.EndTime, validation.Required),
        validation.Field(&booking.StartTime, validation.Required, validation.When(booking.EndTime.After(booking.StartTime))),
        validation.Field(&booking.EndTime, validation.Required, validation.When(booking.EndTime.After(booking.StartTime))),
    )
}
