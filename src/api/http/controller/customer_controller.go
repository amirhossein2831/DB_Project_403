package controller

import (
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Service *services.CustomerService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		Service: services.NewCustomerService(),
	}
}

func (controller *CustomerController) List(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) Get(c *fiber.Ctx) error {
	return nil
}

func (controller *CustomerController) Update(c *fiber.Ctx) error {
	return nil
}

func (controller *CustomerController) Delete(c *fiber.Ctx) error {
	return nil
}
