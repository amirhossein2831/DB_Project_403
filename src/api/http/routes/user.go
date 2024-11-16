package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func UserRoute(router fiber.Router) {
	userController := controller.NewUserController()

	router.Get("/users", userController.List)
}
