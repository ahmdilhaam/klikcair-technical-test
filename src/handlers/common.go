package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
		"code":    fiber.StatusNotFound,
		"status":  "not_found",
		"message": "page not found",
		"data":    nil,
	})
}
