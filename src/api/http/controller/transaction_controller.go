package controller

import (
	"DB_Project/src/api/http/request/transaction"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var TransactionNotFound = errors.New("transaction not found")
var TransactionFieldShouldBeUnique = errors.New("transaction field should be unique: ")
var TransactionRelationNotValid = errors.New("there is no record found for given fk relation in transaction: ")
var TransactionIdNotSet = errors.New("transaction id should be set")

type TransactionController struct {
	Service *services.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		Service: services.NewTransactionService(),
	}
}

func (controller *TransactionController) List(c fiber.Ctx) error {
	sourceId := fiber.Query[int](c, "source_id")

	ctx := context.WithValue(context.Background(), "source_id", sourceId)
	transactions, err := controller.Service.GetTransactions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})
}

func (controller *TransactionController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(TransactionIdNotSet.Error())
	}

	res, err := controller.Service.GetTransaction(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(TransactionNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *TransactionController) Create(c fiber.Ctx) error {
	req := new(transaction.CreateTransactionRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateTransaction(req)
	if err != nil {
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(TransactionFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		if utils.IsErrorCode(err, "23503") {
			return c.Status(fiber.StatusNotFound).SendString(TransactionRelationNotValid.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *TransactionController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(TransactionIdNotSet.Error())
	}

	req := new(transaction.UpdateTransactionRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateTransaction(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *TransactionController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(TransactionIdNotSet.Error())
	}

	err := controller.Service.DeleteTransaction(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(TransactionNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
