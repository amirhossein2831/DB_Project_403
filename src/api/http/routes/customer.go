package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func CustomerRoute(router fiber.Router) {
	customerController := controller.NewCustomerController()

	router.Get("/customers", customerController.List)
}
