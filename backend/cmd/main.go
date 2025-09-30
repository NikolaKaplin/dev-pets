package main

import (
	"dev-pets-backend/internal/config"
	"dev-pets-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var corsConf = cors.Config{
	AllowMethods:     "GET,POST,PUT,DELETE",
	AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-Token",
	ExposeHeaders:    "Set-Cookie",
	AllowOrigins: "*",
	MaxAge:           300,
}

func main() {
	app := fiber.New()

	config.Load()
	
	app.Use(cors.New(corsConf))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(("Hello ебень"))
	})

	app.Get("/google_login", handlers.GoogleLogin)
	app.Get("/google_callback", handlers.GoogleCallback)

	app.Listen(":" + config.ServerConfig().PORT)
}
