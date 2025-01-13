package controller

import (
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type ViewController struct {
	service *services.ViewService
}

func NewViewController() *ViewController {
	return &ViewController{
		service: services.NewViewService(),
	}
}

func (controller *ViewController) CustomerAccounts(c fiber.Ctx) error {
	customers, err := controller.service.GetCustomerAccountsView()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"customers_accounts": customers,
	})
}

func (controller *ViewController) BankTransactions(c fiber.Ctx) error {
	customers, err := controller.service.GetBankTransactionView()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bank_transactions": customers,
	})
}

func (controller *ViewController) BankMembers(c fiber.Ctx) error {
	customers, err := controller.service.GetBankMemberView()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bank_members": customers,
	})
}
