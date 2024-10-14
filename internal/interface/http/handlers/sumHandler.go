package handlers

import (
	"fib/internal/models"

	"github.com/gofiber/fiber/v2"
)

func SumHandler(c *fiber.Ctx) error {
	var r models.Request
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	sum := 0
	for _, num := range r.Nums {
		sum += num
	}

	return c.JSON(fiber.Map{"sum": sum})
}
