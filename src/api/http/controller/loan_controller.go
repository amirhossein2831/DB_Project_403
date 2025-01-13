package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/loan"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"context"
	"github.com/gofiber/fiber/v3"
)

type LoanController struct {
	Service          *services.LoanService
	ExceptionHandler exception.Exception
}

func NewLoanController() *LoanController {
	return &LoanController{
		Service:          services.NewLoanService(),
		ExceptionHandler: exception.NewLoanExceptions(),
	}
}

func (controller *LoanController) List(c fiber.Ctx) error {
	status := fiber.Query[string](c, "status")

	ctx := context.WithValue(context.Background(), "status", status)
	loans, err := controller.Service.GetLoans(ctx)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) ListWithMinInstallmentsPaid(c fiber.Ctx) error {
	loans, err := controller.Service.GetLoansWithMinInstallmentsPaid()
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.LoanIdNotSet.Error())
	}

	res, err := controller.Service.GetLoan(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
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
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *LoanController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.LoanIdNotSet.Error())
	}

	req := new(loan.UpdateLoanRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateLoan(req, id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusOK).Send([]byte{})
}

func (controller *LoanController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.LoanIdNotSet.Error())
	}

	err := controller.Service.DeleteLoan(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
