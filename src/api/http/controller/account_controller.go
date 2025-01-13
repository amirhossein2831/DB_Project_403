package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/account"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"context"
	"github.com/gofiber/fiber/v3"
)

type AccountController struct {
	Service          *services.AccountService
	ExceptionHandler exception.Exception
}

func NewAccountController() *AccountController {
	return &AccountController{
		Service:          services.NewAccountService(),
		ExceptionHandler: exception.NewAccountExceptions(),
	}
}

func (controller *AccountController) List(c fiber.Ctx) error {
	status := fiber.Query[string](c, "status")
	minAmount := fiber.Query[float64](c, "min_amount")

	ctx := context.WithValue(context.Background(), "status", status)
	ctx = context.WithValue(ctx, "min_amount", minAmount)
	accounts, err := controller.Service.GetAccounts(ctx)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accounts": accounts,
	})
}

func (controller *AccountController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.AccountIdNotSet.Error())
	}

	res, err := controller.Service.GetAccount(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
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
		return controller.ExceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *AccountController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.AccountIdNotSet.Error())
	}

	req := new(account.UpdateAccountRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateAccount(req, id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *AccountController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.AccountIdNotSet.Error())
	}

	err := controller.Service.DeleteAccount(id)
	if err != nil {
		return controller.ExceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
