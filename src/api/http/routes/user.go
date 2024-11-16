package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func UserRoute(router fiber.Router) {
	router.Get("/users", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("%s", "hello")
		return c.SendString(msg) // => ✋ register
	})
}
