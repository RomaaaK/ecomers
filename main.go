package main

import (
	"example/ecomers/handlers"
	"example/ecomers/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {

	// add .env setting file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erorr loading .env file")
	}

	engine := html.New("./views", ".html")

	services.InitDataBase()

	app := fiber.New(fiber.Config{
		Views:        engine,
		ViewsLayout:  "layouts/main",
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Static("/", "./public")

	Routes(app)

	app.Listen(":3000")
}
