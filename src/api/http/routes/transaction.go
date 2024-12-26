package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func TransactionRoute(router fiber.Router) {
	transactionController := controller.NewTransactionController()

	router.Get("/transactions", transactionController.List)
	router.Get("/transactions/:id", transactionController.Get)
	router.Post("/transactions", transactionController.Create)
	router.Patch("/transactions/:id", transactionController.Update)
	router.Delete("/transactions/:id", transactionController.Delete)
}
