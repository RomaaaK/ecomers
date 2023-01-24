package handlers

import (
	"example/ecomers/helpers"

	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {

	return c.Render("blog/index", helpers.PreloadMainLayoutData(fiber.Map{}))
}

func BlogById(c *fiber.Ctx) error {

	return c.Render("blog/blog-single", helpers.PreloadMainLayoutData(fiber.Map{}))
}
