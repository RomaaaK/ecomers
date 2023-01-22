package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	var brands []models.Brand
	var categories []models.Category
	var product models.Product

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)
	services.DB.First(&product, id)

	if product.ID == 0 {
		return c.SendStatus(404)
	}

	var productBrand models.Brand

	for _, b := range brands {
		if b.ID == *product.BrandID {
			productBrand = b
			break
		}
	}

	return c.Render("product-details", fiber.Map{
		"Brands":       brands,
		"Categories":   categories,
		"Product":      product,
		"ProductBrand": productBrand,
	})
}

func ProductByCategory(c *fiber.Ctx) error {
	var brands []models.Brand
	var categories []models.Category

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	var products []models.Product

	services.DB.Where("category_id = ?", id).Find(&products)

	return c.Render("index", fiber.Map{
		"Brands":     brands,
		"Categories": categories,
		"Products":   products,
	})
}

func ProductByBrand(c *fiber.Ctx) error {
	var brands []models.Brand
	var categories []models.Category

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	var products []models.Product

	services.DB.Where("brand_id = ?", id).Find(&products)

	return c.Render("index", fiber.Map{
		"Brands":     brands,
		"Categories": categories,
		"Products":   products,
	})
}
