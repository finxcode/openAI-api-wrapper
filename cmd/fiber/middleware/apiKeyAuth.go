package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func ApiKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("middleware here")
		return c.Next()
	}
}
