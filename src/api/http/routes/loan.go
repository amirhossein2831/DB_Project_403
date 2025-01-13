package routes

import (
	"DB_Project/src/api/http/controllers"
	"github.com/gofiber/fiber/v3"
)

func LoanRoute(router fiber.Router) {
	loanController := controllers.NewLoanController()

	router.Get("/loans", loanController.List)
	router.Get("/loans-with-min-installments", loanController.ListWithMinInstallmentsPaid)
	router.Get("/loans/:id", loanController.Get)
	router.Post("/loans", loanController.Create)
	router.Patch("/loans/:id", loanController.Update)
	router.Delete("/loans/:id", loanController.Delete)
}
