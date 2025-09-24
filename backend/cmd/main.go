package main

import (
	"dev-pets-backend/internal/config"
	"dev-pets-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Load()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(("Hello ебень"))
	})

	app.Get("/google_login", handlers.GoogleLogin)
	app.Get("/google_callback", handlers.GoogleCallback)

	app.Listen(":" + config.ServerConfig().PORT)
}
