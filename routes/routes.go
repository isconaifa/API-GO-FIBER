package routes

import (
"github.com/gofiber/fiber/v2"
"fiber/contollers"
)

func SetupRoutes(app *fiber.App) {

	renderer := contollers.Renderer()
	
	app.Get("/", contollers.HomePage)
	app.Get("/users", contollers.GetUsers)
	app.Get("/users/:id", contollers.GetUser)
	app.Post("/users", contollers.AddUser)
	app.Put("/users/:id", contollers.UpdateUser)
	app.Delete("/users/:id", contollers.DeleteUser)
	app.Get("/:page", contollers.DynamicPageHandler(renderer))
}