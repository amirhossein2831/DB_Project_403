package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func ViewsRoute(router fiber.Router) {
	viewController := controller.NewViewController()

	router.Get("/customer_accounts", viewController.CustomerAccounts)
	router.Get("/bank_transactions", viewController.BankTransactions)
	router.Get("/bank_members", viewController.BankMembers)
}
