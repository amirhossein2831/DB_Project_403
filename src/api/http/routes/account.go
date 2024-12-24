package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func AccountRoute(router fiber.Router) {
	accountController := controller.NewAccountController()

	router.Get("/accounts", accountController.List)
	router.Get("/accounts/:id", accountController.Get)
	router.Post("/accounts", accountController.Create)
	router.Patch("/accounts/:id", accountController.Update)
	router.Delete("/accounts/:id", accountController.Delete)
}