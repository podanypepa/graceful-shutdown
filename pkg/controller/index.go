// Package controller ...
package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type resIndex struct {
	TS time.Time `json:""`
}

// Index controller
// GET /
func Index(c *fiber.Ctx) error {
	res := resIndex{
		TS: time.Now(),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
