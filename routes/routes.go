package routes

import (
	"lawas-go/db"
	"lawas-go/models"
	"lawas-go/pages"
	"lawas-go/utils"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {

	// api := app.Group("/api/v1")
	app.Get("/a", func(c *fiber.Ctx) error {
		var u models.User
		db.MySql.Find(&u)
		return c.Status(fiber.StatusOK).JSON(u)
	})
}

func WebRoutes(app *fiber.App) {

	// api := app.Group("/api/v1")
	app.Get("/", func(c *fiber.Ctx) error {
		// return utils.Render(c, components.Layout())
		return utils.Render(c, pages.Frontpage())
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	var u models.User
	// 	db.MySql.Find(&u)
	// 	return c.Status(fiber.StatusOK).JSON(u)
	// })
}
