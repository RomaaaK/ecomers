package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/", "./public")

	Routes(app)

	app.Use(func(c *fiber.Ctx) error {
		c.Context().SetStatusCode(404)
		return c.Render("404", fiber.Map{}, "")
	})

	app.Listen(":3000")
}
