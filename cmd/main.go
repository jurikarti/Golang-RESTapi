package main

import (
	"booking-api/config"
	"booking-api/routes"
	"log"

	_ "booking-api/docs" // Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к базе данных
	config.ConnectDatabase()

	// Создание нового Gin маршрутизатора
	r := gin.Default()

	// Регистрация маршрутов
	routes.RegisterRoutes(r)

	// Подключение Swagger на уникальном маршруте
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера на порту 8080
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
