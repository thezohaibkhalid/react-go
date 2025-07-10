package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //if the valud is false do not add it to the response,
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

// GetTodos handles GET /api/todos
func GetTodos(c *fiber.Ctx) error {
	todos := []Todo{}

	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		fmt.Println("Error fetching todos:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todos",
			"details": err.Error(),
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

func CreateTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body"})
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Todo body cannot be empty",
		})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}

	// Set the inserted ID to the todo
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	fmt.Println(objectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}
	filter := bson.M{"_id": objectID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete todo",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Todo deleted successfully",
		"deletedCount": result.DeletedCount,
	})
}