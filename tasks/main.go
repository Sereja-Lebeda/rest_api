package main

import (
	"log"
	"rest_api/tasks/database"
	"rest_api/tasks/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Инициализируем базу данных
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	// Создаём новое приложение Fiber
	app := fiber.New(fiber.Config{
		Prefork: true, // используем предварительное форкование для увеличения производительности
	})

	// Регистрация маршрутов
	routes.RegisterProductRoutes(app)

	// Запуск сервера
	log.Fatal(app.Listen(":8080"))
}
