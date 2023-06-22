// Package rest ...
package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/podanypepa/graceful-shutdown/pkg/controller"
)

func router(app *fiber.App) {
	app.Get("/", controller.Index)
}
