package exceptions

import "github.com/gofiber/fiber/v3"

type Exception interface {
	Handle(err error, c fiber.Ctx) error
}
