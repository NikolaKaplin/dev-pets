package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	App *fiber.App
}

var corsConf = cors.Config{
	AllowMethods:     "GET,POST,PUT,DELETE",
	AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-Token",
	ExposeHeaders:    "Set-Cookie",
	AllowCredentials: true,
	MaxAge:           300,
}

var limitierConf = limiter.Config{
	Max:        500,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.IP()
	},
}

func (s *Server) mountMiddlewares() {
	s.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	s.App.Use(logger.New())
	s.App.Use(func(c *fiber.Ctx) error {
		c.Context().Conn().SetReadDeadline(time.Now().Add(1 * time.Minute))
		return c.Next()
	})
	s.App.Use(cors.New(corsConf))
	s.App.Use(limiter.New(limitierConf))
}

func (s *Server) mountHandlers() {

}
