package controllers

import (
	"booking-api/models"
	"booking-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Create User"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Error creating user"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	if err := repository.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// DeleteUser godoc
// @Summary Delete a user by username
// @Description Delete a user by username and remove associated bookings
// @Tags users
// @Param username path string true "Username"
// @Success 204
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Error deleting user or bookings"
// @Router /users/{username} [delete]
func DeleteUser(c *gin.Context) {
	username := c.Param("username")

	user, err := repository.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := repository.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	if err := repository.DeleteBookingsByUserID(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user's bookings"})
		return
	}

	c.Status(http.StatusNoContent)
}
