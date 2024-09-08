package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "booking-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=postgres dbname=booking port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to connect to database!")
    }

    database.AutoMigrate(&models.User{}, &models.Booking{})

    DB = database
}
