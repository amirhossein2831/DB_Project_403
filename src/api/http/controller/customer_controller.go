package controller

import (
	"DB_Project/src/api/http/request/customer"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var CustomerNotFound = errors.New("customer not found")
var CustomerFieldShouldBeUnique = errors.New("customer field should be unique: ")
var CustomerRelationNotValid = errors.New("there is no record found for given fk relation in customer: ")
var CustomerIdNotSet = errors.New("customer id should be set")
var CustomerHasActiveLoan = errors.New("customer has active loan")

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

func (controller *CustomerController) ListWithFullName(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomersWithFullName()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndTotalAmount(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomersWithFullNameAndTotalAmount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithFullNameAndAccountNumber(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithFullNameAndAccountNumber()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostLoan(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithMostLoan()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithInstallmentsPenalty(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithInstallmentsPenalty()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) ListWithMostAmount(c fiber.Ctx) error {
	customers, err := controller.Service.GetCustomerWithMostAmount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers": customers,
	})
}

func (controller *CustomerController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(CustomerIdNotSet.Error())
	}

	res, err := controller.Service.GetCustomer(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(CustomerNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
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
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(CustomerFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		if utils.IsErrorCode(err, "23503") {
			return c.Status(fiber.StatusNotFound).SendString(CustomerRelationNotValid.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(CustomerIdNotSet.Error())
	}

	req := new(customer.UpdateCustomerRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateCustomer(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *CustomerController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(CustomerIdNotSet.Error())
	}

	err := controller.Service.DeleteCustomer(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(CustomerNotFound.Error())
		}
		if utils.IsErrorCode(err, "P0001") {
			return c.Status(fiber.StatusUnprocessableEntity).SendString(CustomerHasActiveLoan.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
