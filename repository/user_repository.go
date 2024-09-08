package repository

import (
	"booking-api/models"
	"booking-api/config"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func DeleteUser(user *models.User) error {
	return config.DB.Delete(user).Error
}


