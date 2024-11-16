package controller

import (
	"github.com/gofiber/fiber/v3"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) List(c fiber.Ctx) error {
	return c.SendString("think this is all users")
}
