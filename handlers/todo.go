package handlers

import (
	"todo-app/config"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	config.Database.Find(&todos)
	return c.Status(200).JSON(todos)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	result := config.Database.Find(&todo, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&todo)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.Database.Create(&todo)
	return c.Status(201).JSON(todo)
}

func UpdateTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	config.Database.Model(&models.Todo{}).Where("id = ?", id).Updates(todo)
	return c.Status(200).JSON(todo)
}

func DeleteTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	result := config.Database.Delete(&todo, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
