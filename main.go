package main

import (
	"log"
	"net"
	"tawarin-go/config"
	"tawarin-go/db"
	"tawarin-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.LoadEnv()
	db.ConnectDB()

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Use(logger.New())
	routes.ApiRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	ln, _ := net.Listen("tcp", ":7632")
	log.Fatal(app.Listener(ln))

}
