package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/customer"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type CustomerController struct {
	Service          *services.CustomerService
	ExceptionHandler exception.Exception
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		Service:          services.NewCustomerService(),
		ExceptionHandler: exception.NewCustomerExceptions(),
	}
}

func (controller *CustomerController) List(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomers()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullName(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomersWithFullName()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndTotalAmount(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomersWithFullNameAndTotalAmount()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndAccountNumber(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithFullNameAndAccountNumber()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostLoan(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithMostLoan()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithInstallmentsPenalty(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithInstallmentsPenalty()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostAmount(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithMostAmount()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.CustomerIdNotSet.Error())
	}

	res, err := controller.Service.GetCustomer(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *CustomerController) Create(c fiber.Ctx) error {
	req := new(customer.CreateCustomerRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateCustomer(req)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.CustomerIdNotSet.Error())
	}

	req := new(customer.UpdateCustomerRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateCustomer(req, id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.CustomerIdNotSet.Error())
	}

	err := controller.Service.DeleteCustomer(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
