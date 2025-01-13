package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/employee"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type EmployeeController struct {
	service          *services.EmployeeService
	exceptionHandler exception.Exception
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		service:          services.NewEmployeeService(),
		exceptionHandler: exception.NewEmployeeExceptions(),
	}
}

func (controller *EmployeeController) List(c fiber.Ctx) error {
	employees, err := controller.service.GetEmployees()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"employees": employees,
	})
}

func (controller *EmployeeController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.EmployeeIdNotSet.Error())
	}

	res, err := controller.service.GetEmployee(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *EmployeeController) Create(c fiber.Ctx) error {
	req := new(employee.CreateEmployeeRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateEmployee(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *EmployeeController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.EmployeeIdNotSet.Error())
	}

	req := new(employee.UpdateEmployeeRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.UpdateEmployee(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *EmployeeController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.EmployeeIdNotSet.Error())
	}

	err := controller.service.DeleteEmployee(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
