package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Task represents the structure of a task in the database.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// CreateTask handles the creation of a new task.
func CreateTask(c *fiber.Ctx) error {
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Logic to save the task in the database goes here
	return c.Status(http.StatusCreated).JSON(task)
}

// GetTasks handles fetching all tasks.
func GetTasks(c *fiber.Ctx) error {
	// Logic to retrieve tasks from the database goes here
	var tasks []Task
	return c.Status(http.StatusOK).JSON(tasks)
}

// UpdateTask handles updating an existing task.
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	// Logic to update the task in the database goes here
	return c.Status(http.StatusOK).JSON(task)
}

// DeleteTask handles deleting a task.
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	// Logic to delete the task from the database goes here
	return c.Status(http.StatusNoContent).SendString("")
}