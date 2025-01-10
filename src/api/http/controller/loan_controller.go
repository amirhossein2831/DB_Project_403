package controller

import (
	"DB_Project/src/api/http/request/loan"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var LoanNotFound = errors.New("loan not found")
var LoanFieldShouldBeUnique = errors.New("loan field should be unique: ")
var LoanRelationNotValid = errors.New("there is no record found for given fk relation in loan: ")
var LoanIdNotSet = errors.New("loan id should be set")

type LoanController struct {
	Service *services.LoanService
}

func NewLoanController() *LoanController {
	return &LoanController{
		Service: services.NewLoanService(),
	}
}

func (controller *LoanController) List(c fiber.Ctx) error {
	status := fiber.Query[string](c, "status")

	ctx := context.WithValue(context.Background(), "status", status)
	loans, err := controller.Service.GetLoans(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) ListWithMinInstallmentsPaid(c fiber.Ctx) error {
	loans, err := controller.Service.GetLoansWithMinInstallmentsPaid()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(LoanIdNotSet.Error())
	}

	res, err := controller.Service.GetLoan(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(LoanNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *LoanController) Create(c fiber.Ctx) error {
	req := new(loan.CreateLoanRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateLoan(req)
	if err != nil {
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(LoanFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		if utils.IsErrorCode(err, "23503") {
			return c.Status(fiber.StatusNotFound).SendString(LoanRelationNotValid.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *LoanController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(LoanIdNotSet.Error())
	}

	req := new(loan.UpdateLoanRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateLoan(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).Send([]byte{})
}

func (controller *LoanController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(LoanIdNotSet.Error())
	}

	err := controller.Service.DeleteLoan(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(LoanNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
