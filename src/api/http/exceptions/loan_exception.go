package exceptions

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	LoanNotFound            = errors.New("loan not found with given id")
	LoanFieldShouldBeUnique = errors.New("loan field should be unique: ")
	LoanRelationNotValid    = errors.New("there is no record found for given fk relation in loan: ")
	LoanIdNotSet            = errors.New("loan id should be set")
)

type LoanExceptions struct{}

func NewLoanExceptions() *LoanExceptions {
	return &LoanExceptions{}
}

func (exceptions *LoanExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(LoanNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(LoanFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(LoanRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
