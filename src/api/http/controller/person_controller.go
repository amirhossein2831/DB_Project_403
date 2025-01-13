package controller

import (
	"github.com/gofiber/fiber/v3"
)

type PersonController struct {
}

func NewPersonController() *PersonController {
	return &PersonController{}
}

func (controller *PersonController) List(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"": "",
	})
}

func (controller *PersonController) Get(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "": "",
    })
}

func (controller *PersonController) Create(c fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *PersonController) Update(c fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *PersonController) Delete(c fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
