package repository

import (
	"booking-api/models"
	"booking-api/config"
	"errors"
)

func CreateUser(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	// Дополнительная проверка на уникальность имени пользователя
	existingUser, _ := GetUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	return config.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(user *models.User) error {
	return config.DB.Delete(user).Error
}
