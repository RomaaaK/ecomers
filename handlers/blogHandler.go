package handlers

import (
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {

	return c.Render("blog/index", fiber.Map{
		"Posts": services.GetPosts(),
	})
}

func BlogById(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	post, err := services.GetPostById(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.Render("blog/post", fiber.Map{
		"Post": post,
	})
}
