package controllers

import (
	"DB_Project/src/api/http/exceptions"
	"DB_Project/src/api/http/requests/loan"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"context"
	"github.com/gofiber/fiber/v3"
)

type LoanController struct {
	service          *services.LoanService
	exceptionHandler exceptions.Exception
}

func NewLoanController() *LoanController {
	return &LoanController{
		service:          services.NewLoanService(),
		exceptionHandler: exceptions.NewLoanExceptions(),
	}
}

func (controller *LoanController) List(c fiber.Ctx) error {
	status := fiber.Query[string](c, "status")

	ctx := context.WithValue(context.Background(), "status", status)
	loans, err := controller.service.GetLoans(ctx)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) ListWithMinInstallmentsPaid(c fiber.Ctx) error {
	loans, err := controller.service.GetLoansWithMinInstallmentsPaid()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"loans": loans,
	})
}

func (controller *LoanController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.LoanIdNotSet.Error())
	}

	res, err := controller.service.GetLoan(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *LoanController) Create(c fiber.Ctx) error {
	req := new(loan.CreateLoanRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateLoan(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *LoanController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.LoanIdNotSet.Error())
	}

	req := new(loan.UpdateLoanRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.UpdateLoan(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusOK).Send([]byte{})
}

func (controller *LoanController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exceptions.LoanIdNotSet.Error())
	}

	err := controller.service.DeleteLoan(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
