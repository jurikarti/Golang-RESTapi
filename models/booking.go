package models

import "time"

type Booking struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"user_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}