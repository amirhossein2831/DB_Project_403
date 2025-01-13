package controllers

import (
	"DB_Project/src/api/http/exceptions"
	"DB_Project/src/api/http/requests/customer"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	service          *services.CustomerService
	exceptionHandler exceptions.Exception
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		service:          services.NewCustomerService(),
		exceptionHandler: exceptions.NewCustomerExceptions(),
	}
}

func (controller *CustomerController) List(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomers()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullName(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomersWithFullName()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndTotalAmount(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomersWithFullNameAndTotalAmount()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndAccountNumber(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomerWithFullNameAndAccountNumber()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostLoan(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomerWithMostLoan()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithInstallmentsPenalty(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomerWithInstallmentsPenalty()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostAmount(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomerWithMostAmount()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.CustomerIdNotSet.Error())
	}

	res, err := controller.service.GetCustomer(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *CustomerController) Create(c fiber.Ctx) error {
	req := new(customer.CreateCustomerRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateCustomer(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.CustomerIdNotSet.Error())
	}

	req := new(customer.UpdateCustomerRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.UpdateCustomer(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.CustomerIdNotSet.Error())
	}

	err := controller.service.DeleteCustomer(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
