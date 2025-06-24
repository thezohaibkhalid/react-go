package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	// In-memory storage for todos
	todos := []Todo{}
	var idCounter int = 1

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

		// Assign a unique ID
		todo.ID = idCounter
		idCounter++

		// Append to todos slice
		todos = append(todos, *todo)

		// var x int = 5   //store in memory address 0x000012 and value is 5
		// var p *int = &x // p is a pointer to memory address of x, storing the address of x

		// fmt.Println("Value of x:", *p)  // prints 5
		// fmt.Println("Address of x:", p) // prints memory address of x

		return c.Status(201).JSON(todo)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//Update a todo
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

	log.Fatal(app.Listen(":" + port))
}
