package repository

import (
    "booking-api/config"
    "booking-api/models"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func clearDatabase() {
    config.DB.Exec("DELETE FROM bookings")
    config.DB.Exec("DELETE FROM users")
}

func TestCreateBooking(t *testing.T) {
    config.ConnectDatabase() // Подключаемся к тестовой базе данных
    defer config.CloseDatabase()

    clearDatabase() // Очищаем базу данных перед тестом

    user := &models.User{
        Username: "testuser",
        Password: "testpassword",
    }
    err := user.HashPassword(user.Password)
    assert.NoError(t, err)

    // Проверяем валидацию перед созданием пользователя
    err = user.Validate()
    assert.NoError(t, err)

    err = CreateUser(user)
    assert.NoError(t, err)

    // Перезагружаем пользователя, чтобы получить правильный ID
    var createdUser models.User
    err = config.DB.Where("username = ?", user.Username).First(&createdUser).Error
    assert.NoError(t, err)

    booking := &models.Booking{
        UserID:    createdUser.ID, // Используем ID созданного пользователя
        StartTime: time.Now().Add(1 * time.Hour),
        EndTime:   time.Now().Add(2 * time.Hour),
    }

    // Проверяем валидацию перед созданием бронирования
    err = booking.Validate()
    assert.NoError(t, err)

    err = CreateBooking(booking)
    assert.NoError(t, err)

    // Проверяем, что бронирование было создано
    var retrievedBooking models.Booking
    err = config.DB.First(&retrievedBooking, booking.ID).Error
    assert.NoError(t, err)
    assert.Equal(t, booking.UserID, retrievedBooking.UserID)
}

func TestDeleteBookingsByUserID(t *testing.T) {
    config.ConnectDatabase() // Подключаемся к тестовой базе данных
    defer config.CloseDatabase()

    clearDatabase() // Очищаем базу данных перед тестом

    user := &models.User{
        Username: "testuser",
        Password: "testpassword",
    }
    err := user.HashPassword(user.Password)
    assert.NoError(t, err)

    // Проверяем валидацию перед созданием пользователя
    err = user.Validate()
    assert.NoError(t, err)

    err = CreateUser(user)
    assert.NoError(t, err)

    // Перезагружаем пользователя, чтобы получить правильный ID
    var createdUser models.User
    err = config.DB.Where("username = ?", user.Username).First(&createdUser).Error
    assert.NoError(t, err)

    booking := &models.Booking{
        UserID:    createdUser.ID, // Используем ID созданного пользователя
        StartTime: time.Now().Add(1 * time.Hour),
        EndTime:   time.Now().Add(2 * time.Hour),
    }

    // Проверяем валидацию перед созданием бронирования
    err = booking.Validate()
    assert.NoError(t, err)

    err = CreateBooking(booking)
    assert.NoError(t, err)

    err = DeleteBookingsByUserID(createdUser.ID)
    assert.NoError(t, err)

    // Проверяем, что бронирование было удалено
    var bookings []models.Booking
    err = config.DB.Where("user_id = ?", createdUser.ID).Find(&bookings).Error
    assert.NoError(t, err)
    assert.Empty(t, bookings)
}
