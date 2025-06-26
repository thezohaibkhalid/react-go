package main

import (
	"fmt"
	"log"
	"os"

	""
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos []Todo
var idCounter int = 1

func main() {
	app := fiber.New()

	// GET /api/todos - Retrieve all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// POST /api/todos - Create a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}
		if err := c.BodyParser(&todo); err != nil {
			log.Printf("Body parsing error: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body cannot be empty"})
		}

		todo.ID = idCounter
		idCounter++

		todos = append(todos, todo)

		return c.Status(201).JSON(todo)
	})

	// PATCH /api/todos/:id - Mark todo as completed
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// DELETE /api/todos/:id - Delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"message": "Todo deleted successfully"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(app.Listen(":" + port))
}
