package controller

import (
	"DB_Project/src/api/http/request/employee"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var EmployeeNotFound = errors.New("employee not found")
var EmployeeFieldShouldBeUnique = errors.New("employee field should be unique: ")

type EmployeeController struct {
	Service *services.EmployeeService
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		Service: services.NewEmployeeService(),
	}
}

func (controller *EmployeeController) List(c fiber.Ctx) error {
	employees, err := controller.Service.GetEmployees()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"employees": employees,
	})
}

func (controller *EmployeeController) Get(c fiber.Ctx) error {
	id := c.Params("id")

	res, err := controller.Service.GetEmployee(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(EmployeeNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *EmployeeController) Create(c fiber.Ctx) error {
	req := new(employee.CreateEmployeeRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateEmployee(req)
	if err != nil {
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(EmployeeFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *EmployeeController) Update(c fiber.Ctx) error {
	id := c.Params("id")

	req := new(employee.UpdateEmployeeRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateEmployee(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *EmployeeController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	err := controller.Service.DeleteEmployee(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(EmployeeNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
