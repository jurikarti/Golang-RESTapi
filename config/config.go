package config

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "booking-api/models"
)

var DB *gorm.DB

// ConnectDatabase инициализирует соединение с бд
func ConnectDatabase() {
    // запуск без докера
    dsn := "host=localhost user=postgres password=btS2RU6r dbname=booking port=5432 sslmode=disable TimeZone=Asia/Shanghai"

    //запуск для докера
    //dsn := "host=host.docker.internal user=postgres password=btS2RU6r dbname=booking port=5432 sslmode=disable TimeZone=Asia/Shanghai"

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // AutoMigrate автоматически создаст таблицы, недостающие столбцы и индексы.
    if err := DB.AutoMigrate(&models.User{}, &models.Booking{}); err != nil {
        log.Fatalf("Failed to auto migrate: %v", err)
    }
}

// CloseDatabase закрывает соединение с базой данных
func CloseDatabase() {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Error getting DB instance: %v", err)
    }
    if err := sqlDB.Close(); err != nil {
        log.Fatalf("Error closing database connection: %v", err)
    }
}
