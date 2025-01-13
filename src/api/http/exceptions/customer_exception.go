package exceptions

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	CustomerNotFound            = errors.New("customer not found")
	CustomerFieldShouldBeUnique = errors.New("customer field should be unique: ")
	CustomerRelationNotValid    = errors.New("there is no record found for given fk relation in customer: ")
	CustomerIdNotSet            = errors.New("customer id should be set")
	CustomerHasActiveLoan       = errors.New("customer has active loan")
)

type CustomerExceptions struct{}

func NewCustomerExceptions() *CustomerExceptions {
	return &CustomerExceptions{}
}

func (exceptions *CustomerExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(CustomerNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(CustomerFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(CustomerRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "P0001"):
		return c.Status(fiber.StatusUnprocessableEntity).SendString(CustomerHasActiveLoan.Error())

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
