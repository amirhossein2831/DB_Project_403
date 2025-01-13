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
	service          *services.AccountService
	exceptionHandler exception.Exception
}

func NewAccountController() *AccountController {
	return &AccountController{
		service:          services.NewAccountService(),
		exceptionHandler: exception.NewAccountExceptions(),
	}
}

func (controller *AccountController) List(c fiber.Ctx) error {
	status := fiber.Query[string](c, "status")
	minAmount := fiber.Query[float64](c, "min_amount")

	ctx := context.WithValue(context.Background(), "status", status)
	ctx = context.WithValue(ctx, "min_amount", minAmount)
	accounts, err := controller.service.GetAccounts(ctx)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
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

	res, err := controller.service.GetAccount(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *AccountController) Create(c fiber.Ctx) error {
	req := new(account.CreateAccountRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateAccount(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
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

	err := controller.service.UpdateAccount(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *AccountController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.AccountIdNotSet.Error())
	}

	err := controller.service.DeleteAccount(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
