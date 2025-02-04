package main

import (
	"lawas-go/config"
	"lawas-go/db"
	"lawas-go/routes"
	"log"
	"net"

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

	app.Static("/assets", "./assets")
	app.Static("/lib", "./lib")
	app.Static("/popup", "./popup")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		c.Set("Surrogate-Control", "no-store")
		return c.Next()
	})

	routes.ApiRoutes(app)
	routes.WebRoutes(app)

	listen, _ := net.Listen("tcp", ":7632")
	log.Fatal(app.Listener(listen))

}
