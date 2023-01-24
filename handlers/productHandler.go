package handlers

import (
	"example/ecomers/helpers"
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	product, err := services.GetProductById(id)

	if err != nil {
		return c.SendStatus(404)
	}

	return c.Render("product-details", helpers.PreloadMainLayoutData(fiber.Map{
		"Product": product,
	}))
}

func ProductByCategory(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	return c.Render("index", helpers.PreloadMainLayoutData(fiber.Map{
		"Products": services.GetProductsByCategory(id),
	}))
}

func ProductByBrand(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	return c.Render("index", helpers.PreloadMainLayoutData(fiber.Map{
		"Products": services.GetProductsByBrand(id),
	}))
}
