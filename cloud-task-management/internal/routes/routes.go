package routes

import (
    "github.com/gofiber/fiber/v2"
    "cloud-task-management/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    // Task routes
    api.Post("/tasks", handlers.CreateTask)
    api.Get("/tasks/:id", handlers.GetTask)
    api.Put("/tasks/:id", handlers.UpdateTask)
    api.Delete("/tasks/:id", handlers.DeleteTask)
    api.Get("/tasks", handlers.GetAllTasks)
}