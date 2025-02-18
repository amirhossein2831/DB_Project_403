package routes

import (
	"DB_Project/src/api/http/controllers"
	"github.com/gofiber/fiber/v3"
)

func InstallmentRoute(router fiber.Router) {
	installmentController := controllers.NewInstallmentController()

	router.Get("/installments", installmentController.List)
	router.Get("/installments/:id", installmentController.Get)
	router.Post("/installments", installmentController.Create)
	router.Patch("/installments/:id", installmentController.Update)
	router.Delete("/installments/:id", installmentController.Delete)
}
