package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func InstallmentRoute(router fiber.Router) {
	installmentController := controller.NewInstallmentController()

	router.Get("/installment", installmentController.List)
	router.Get("/installment/:id", installmentController.Get)
	router.Post("/installment", installmentController.Create)
	router.Patch("/installment/:id", installmentController.Update)
	router.Delete("/installment/:id", installmentController.Delete)
}
