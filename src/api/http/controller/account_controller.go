package controller

import (
	"DB_Project/src/api/http/request/account"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var AccountNotFound = errors.New("account not found")
var AccountFieldShouldBeUnique = errors.New("account field should be unique: ")
var AccountIdNotSet = errors.New("account id should be set")

type AccountController struct {
	Service *services.AccountService
}

func NewAccountController() *AccountController {
	return &AccountController{
		Service: services.NewAccountService(),
	}
}

func (controller *AccountController) List(c fiber.Ctx) error {
	accounts, err := controller.Service.GetAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accounts": accounts,
	})
}

func (controller *AccountController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(AccountIdNotSet.Error())
	}
	
	res, err := controller.Service.GetAccount(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(AccountNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *AccountController) Create(c fiber.Ctx) error {
	req := new(account.CreateAccountRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateAccount(req)
	if err != nil {
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(AccountFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *AccountController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(AccountIdNotSet.Error())
	}

	req := new(account.UpdateAccountRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateAccount(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *AccountController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(AccountIdNotSet.Error())
	}

	err := controller.Service.DeleteAccount(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(AccountNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
