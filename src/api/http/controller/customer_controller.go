package controller

import (
	"DB_Project/src/api/http/request/customer"
	"DB_Project/src/models"
	"DB_Project/src/services"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
	"time"
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

func (controller *CustomerController) Create(c fiber.Ctx) error {
	req := &customer.CreateCustomerRequest{
		FirstName:    "amir",
		LastName:     "sda",
		BirthDate:    time.Now(),
		Phone:        "902938485",
		Email:        "amirmemool1ssds2@gmail.cm",
		Address:      "sdfdsf",
		CustomerType: models.IndividualCustomerType,
	}

	err := controller.Service.CreateCustomer(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Update(c *fiber.Ctx) error {
	return nil
}

func (controller *CustomerController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	err := controller.Service.DeleteCustomer(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(CustomerNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
