package repository

import (
    "booking-api/config"
    "booking-api/models"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    config.ConnectDatabase() // Подключаемся к тестовой базе данных
    defer config.CloseDatabase()

    clearDatabase() // Очищаем базу данных перед тестом

    user := &models.User{
        Username: "uniqueuser",
        Password: "testpassword",
    }
    err := user.HashPassword(user.Password)
    assert.NoError(t, err)

    // Проверяем валидацию перед созданием пользователя
    err = user.Validate()
    assert.NoError(t, err)

    err = CreateUser(user)
    assert.NoError(t, err)

    // Проверяем, что пользователь был создан
    var retrievedUser models.User
    err = config.DB.Where("username = ?", user.Username).First(&retrievedUser).Error
    assert.NoError(t, err)
    assert.Equal(t, user.Username, retrievedUser.Username)
}

func TestDeleteUser(t *testing.T) {
    config.ConnectDatabase() // Подключаемся к тестовой базе данных
    defer config.CloseDatabase()

    clearDatabase() // Очищаем базу данных перед тестом

    user := &models.User{
        Username: "uniqueuser",
        Password: "testpassword",
    }
    err := user.HashPassword(user.Password)
    assert.NoError(t, err)

    // Проверяем валидацию перед созданием пользователя
    err = user.Validate()
    assert.NoError(t, err)

    err = CreateUser(user)
    assert.NoError(t, err)

    err = DeleteUser(user)
    assert.NoError(t, err)

    // Проверяем, что пользователь был удален
    var retrievedUser models.User
    err = config.DB.Where("username = ?", user.Username).First(&retrievedUser).Error
    assert.Error(t, err) // Ожидаем ошибку, так как пользователь должен быть удален
}
