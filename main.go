package main

import (
	"log"
	"todo-app/config"
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/api/todos", handlers.GetTodos)
	app.Get("/api/todos/:id", handlers.GetTodoById)
	app.Post("/api/todos", handlers.AddTodo)
	app.Put("/api/todos/:id", handlers.UpdateTodoById)
	app.Delete("/api/todos/:id", handlers.DeleteTodoById)

	log.Fatal(app.Listen(":3000"))
}
