package helpers

import (
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func PreloadMainLayoutData(data fiber.Map) fiber.Map {

	result := fiber.Map{
		"Brands":     services.GetBrands(),
		"Categories": services.GetCategories(),
	}

	for key, value := range data {
		result[key] = value
	}

	return result
}
