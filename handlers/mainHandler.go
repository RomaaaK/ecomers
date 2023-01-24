package handlers

import (
	"example/ecomers/helpers"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	return c.Render("index", helpers.PreloadMainLayoutData(fiber.Map{
		"Slider":   true,
		"Products": services.GetProducts(),
	}))
}
