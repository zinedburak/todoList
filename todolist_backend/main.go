package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todolist_backend/database"
	"todolist_backend/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))


	routes.Setup(app)

	app.Listen(":8000")
}

