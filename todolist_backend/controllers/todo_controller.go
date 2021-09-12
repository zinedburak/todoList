package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"todolist_backend/database"
	"todolist_backend/models"
)

func GetTodos(c *fiber.Ctx) error{
	var todos []models.Todos
	result := database.DB.Find(&todos)
	if (result.Error != nil){
		fmt.Println("There was an error")
		return result.Error
	}
	return c.JSON(todos)
}

func AddTodos(c *fiber.Ctx) error {
	var ctodo map[string]string
	err := c.BodyParser(&ctodo)
	if err != nil{

		return err
	}
	todo := models.Todos{
		Text:   ctodo["text"],
	}
	database.DB.Create(&todo)
	return c.JSON(todo)
}