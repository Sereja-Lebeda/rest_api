package routes

import (
	"rest_api/tasks/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/tasks", handlers.GetTasks)
	api.Post("/tasks", handlers.CreateTask)
	api.Get("/tasks/:id", handlers.GetTask)
	api.Put("/tasks/:id", handlers.UpdateTask)
	api.Delete("/tasks/:id", handlers.DeleteTask)
}
