package exceptions

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	TransactionNotFound            = errors.New("transaction not found")
	TransactionFieldShouldBeUnique = errors.New("transaction field should be unique: ")
	TransactionRelationNotValid    = errors.New("there is no record found for given fk relation in transaction: ")
	TransactionIdNotSet            = errors.New("transaction id should be set")
	NotEnoughAmount                = errors.New("account Amount is not enough")
	DestinationIdNotSet            = errors.New("destination id not set")
)

type TransactionExceptions struct{}

func NewTransactionExceptions() *TransactionExceptions {
	return &TransactionExceptions{}
}

func (exceptions *TransactionExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(TransactionNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(TransactionFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(TransactionRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "P0001"):
		return c.Status(fiber.StatusUnprocessableEntity).SendString(NotEnoughAmount.Error())

	case utils.IsErrorCode(err, "P0002"):
		return c.Status(fiber.StatusBadRequest).SendString(DestinationIdNotSet.Error())

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
