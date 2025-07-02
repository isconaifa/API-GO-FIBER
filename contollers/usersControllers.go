package contollers

import (
	"fiber/model"
	"os"
	"strings"
	"text/template"
	"github.com/gofiber/fiber/v2"
)


func HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Aplicação Go Fiber",
		"Handing": "Olá, Devs!",
		"Message": "Sejam bem-vindos ao Fiber, um poderoso framework do Go ! ",
	})
}

func Renderer() *template.Template{
	return template.Must(template.ParseGlob("views/*.html"))
}

func DynamicPageHandler(_ *template.Template) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.Params("page")

		if strings.HasSuffix(page, ".html") {
			page = strings.TrimSuffix(page, ".html")
		}
		
		if _, err :=os.Stat("views/" + page + ".html"); err == nil {
			return c.Render(page, nil)
		}

		 return c.Status(fiber.StatusNotFound).SendString("Pagina não encontrado!")
	}
   
}

// AddUser add a new user
func AddUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
    c.SendStatus(fiber.StatusCreated)
	return c.JSON(user)
}


// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Getting user " + id)
	
}


// GetUsers get all users
func GetUsers(c *fiber.Ctx) error {
	user := model.User{
		ID: 1,
		Name: "John Doe",
		Email: "j@j.com",
		Password: "123456",
	}
	c.SendStatus(fiber.StatusOK)
	return c.JSON(user)
}


// UpdateUser update a user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Updating user " + id)


}


// DeleteUser delete a user
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Deleting user " + id)
}