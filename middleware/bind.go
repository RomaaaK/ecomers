package middleware

import (
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func BindMainLayoutData(c *fiber.Ctx) error {
	c.Bind(fiber.Map{
		"Brands":     services.GetBrands(),
		"Categories": services.GetCategories(),
	})

	return c.Next()
}
