package exceptions

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	InstallmentNotFound            = errors.New("installment not found")
	InstallmentFieldShouldBeUnique = errors.New("installment field should be unique: ")
	InstallmentRelationNotValid    = errors.New("there is no record found for given fk relation in installment: ")
	InstallmentIdNotSet            = errors.New("installment id should be set")
)

type InstallmentExceptions struct{}

func NewInstallmentExceptions() *InstallmentExceptions {
	return &InstallmentExceptions{}
}

func (exceptions *InstallmentExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(InstallmentNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(InstallmentFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(InstallmentRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "P0001"):
		return c.Status(fiber.StatusUnprocessableEntity).SendString(CustomerHasActiveLoan.Error())

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
