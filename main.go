package main

import (
	"fiber/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New()
    
	engine := html.New("./views", ".html")
	app = fiber.New(fiber.Config{
		Views: engine,
	})
     
    app.Static("/static", "./static")

	app.Use(logginMiddleware)
	fmt.Println("Server is running on port 3000")
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func logginMiddleware(c *fiber.Ctx) error {
	log.Printf("solicitação recebida: %s %s", c.Method(), c.Path())
	return c.Next()
}
