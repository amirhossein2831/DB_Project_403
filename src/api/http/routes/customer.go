package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func CustomerRoute(router fiber.Router) {
	customerController := controller.NewCustomerController()

	router.Get("/customers", customerController.List)
	router.Get("/customers/:id", customerController.Get)
	router.Post("/customers", customerController.Create)
	router.Patch("/customers/:id", customerController.Update)
	router.Delete("/customers/:id", customerController.Delete)
}
