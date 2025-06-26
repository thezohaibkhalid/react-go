package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	ID        int    `json:"_id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}


// GetTodos handles GET /api/todos
func GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todos",
		})
	}
	defer cursor.Close(c.Context())

	for cursor.Next(c.Context()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to decode todo",
			})
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}
