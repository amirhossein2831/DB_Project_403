package exception

import (
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var (
	EmployeeNotFound            = errors.New("employee not found")
	EmployeeFieldShouldBeUnique = errors.New("employee field should be unique: ")
	EmployeeRelationNotValid    = errors.New("there is no record found for given fk relation in employee: ")
	EmployeeIdNotSet            = errors.New("employee id should be set")
)

type EmployeeExceptions struct{}

func NewEmployeeExceptions() *EmployeeExceptions {
	return &EmployeeExceptions{}
}

func (exceptions *EmployeeExceptions) Handle(err error, c fiber.Ctx) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return c.Status(fiber.StatusNotFound).SendString(EmployeeNotFound.Error())

	case utils.IsErrorCode(err, "23505"):
		return c.Status(fiber.StatusConflict).SendString(EmployeeFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))

	case utils.IsErrorCode(err, "23503"):
		return c.Status(fiber.StatusNotFound).SendString(EmployeeRelationNotValid.Error() + utils.GetErrorConstraintName(err))

	default:
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}
