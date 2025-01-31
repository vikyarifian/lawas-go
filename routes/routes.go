package routes

import (
	"tawarin-go/db"
	"tawarin-go/models"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {

	// api := app.Group("/api/v1")
	app.Get("/", func(c *fiber.Ctx) error {
		var u models.User
		db.MySql.Find(&u)
		return c.Status(fiber.StatusOK).JSON(u)
	})
}

func WebRoutes(app *fiber.App) {

	// api := app.Group("/api/v1")
	app.Get("/a", func(c *fiber.Ctx) error {
		var u models.User
		db.MySql.Find(&u)
		return c.Status(fiber.StatusOK).JSON(u)
	})
}
