package exception

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	AccountNotFound            = errors.New("account not found")
	AccountFieldShouldBeUnique = errors.New("account field should be unique: ")
	AccountRelationNotValid    = errors.New("there is no record found for given fk relation in account: ")
	AccountIdNotSet            = errors.New("account id should be set")
)

type AccountExceptions struct{}

func NewAccountExceptions() *AccountExceptions {
	return &AccountExceptions{}
}

func (exceptions *AccountExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(AccountNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(AccountFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(AccountRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
