package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func CustomerRoute(router fiber.Router) {
	customerController := controller.NewCustomerController()

	router.Get("/customers", customerController.List)
	router.Get("/customers-full-name", customerController.ListWithFullName)
	router.Get("/customers-full-name-with-total-amount", customerController.ListWithFullNameAndTotalAmount)
	router.Get("/customers-full-name-with-account-number", customerController.ListWithFullNameAndAccountNumber)
	router.Get("/customers/:id", customerController.Get)
	router.Post("/customers", customerController.Create)
	router.Patch("/customers/:id", customerController.Update)
	router.Delete("/customers/:id", customerController.Delete)
}
