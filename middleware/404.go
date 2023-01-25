package middleware

import "github.com/gofiber/fiber/v2"

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return c.Render("404", fiber.Map{})
}
