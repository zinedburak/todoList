package routes


import (
	"github.com/gofiber/fiber/v2"
	"todolist_backend/controllers"
)

func Setup(app *fiber.App)  {


	app.Get("/api/get_todos",controllers.GetTodos)
	app.Post("/api/add_todos",controllers.AddTodos)

}