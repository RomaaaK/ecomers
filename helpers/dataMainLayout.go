package helpers

import (
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func PreloadMainLayoutData(data fiber.Map) fiber.Map {

	var brands []models.Brand
	var categories []models.Category

	services.DB.Model(&models.Category{}).Where("category_id IS NULL").Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)

	result := fiber.Map{
		"Brands":     brands,
		"Categories": categories,
	}

	for key, value := range data {
		result[key] = value
	}

	return result
}
