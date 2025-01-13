package routes

import (
	"DB_Project/src/api/http/controllers"
	"github.com/gofiber/fiber/v3"
)

func ViewsRoute(router fiber.Router) {
	viewController := controllers.NewViewController()

	router.Get("/customer_accounts", viewController.CustomerAccounts)
	router.Get("/bank_transactions", viewController.BankTransactions)
	router.Get("/bank_members", viewController.BankMembers)
}
