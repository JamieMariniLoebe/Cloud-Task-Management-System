package main

import (
    "github.com/gofiber/fiber/v2"
    "cloud-task-management/internal/routes"
)

func main() {
    app := fiber.New()

    // Set up middleware here if needed

    // Set up routes
    routes.SetupRoutes(app)

    // Start the server
    app.Listen(":3000")
}