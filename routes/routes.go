package routes

import (
	"booking-api/controllers"
	"github.com/gin-gonic/gin"
	_ "booking-api/docs" // Swagger docs
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// RegisterRoutes регистрирует маршруты с использованием Gin
func RegisterRoutes(r *gin.Engine) {
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:username", controllers.DeleteUser)

	// bookings
	r.POST("/bookings", controllers.CreateBooking)
	r.GET("/bookings/:user_id", controllers.GetBookings)
}
