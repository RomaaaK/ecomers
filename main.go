package main

import (
	"example/ecomers/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erorr loading .env file")
	}

	engine := html.New("./views", ".html")

	services.InitDataBase()

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/", "./public")

	Routes(app)

	app.Listen(":3000")
}
