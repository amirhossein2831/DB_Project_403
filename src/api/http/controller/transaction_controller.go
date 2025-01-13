package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/transaction"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"context"
	"github.com/gofiber/fiber/v3"
)

type TransactionController struct {
	service          *services.TransactionService
	exceptionHandler exception.Exception
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		service:          services.NewTransactionService(),
		exceptionHandler: exception.NewTransactionExceptions(),
	}
}

func (controller *TransactionController) List(c fiber.Ctx) error {
	sourceId := fiber.Query[int](c, "source_id")

	ctx := context.WithValue(context.Background(), "source_id", sourceId)
	transactions, err := controller.service.GetTransactions(ctx)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})
}

func (controller *TransactionController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.TransactionIdNotSet.Error())
	}

	res, err := controller.service.GetTransaction(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *TransactionController) Create(c fiber.Ctx) error {
	req := new(transaction.CreateTransactionRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateTransaction(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *TransactionController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.TransactionIdNotSet.Error())
	}

	req := new(transaction.UpdateTransactionRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.UpdateTransaction(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *TransactionController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.TransactionIdNotSet.Error())
	}

	err := controller.service.DeleteTransaction(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
