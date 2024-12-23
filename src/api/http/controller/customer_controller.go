package controller

import (
	"DB_Project/src/services"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var CustomerNotFound = errors.New("customer not found")

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

func (controller *CustomerController) Get(c fiber.Ctx) error {
	id := c.Params("id")

	customer, err := controller.Service.GetCustomer(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(CustomerNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (controller *CustomerController) Update(c *fiber.Ctx) error {
	return nil
}

func (controller *CustomerController) Delete(c *fiber.Ctx) error {
	return nil
}
