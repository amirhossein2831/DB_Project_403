package routes

import (
	"DB_Project/src/api/http/controller"
	"github.com/gofiber/fiber/v3"
)

func EmployeeRoute(router fiber.Router) {
	employeeController := controller.NewEmployeeController()

	router.Get("/employees", employeeController.List)
	router.Get("/employees/:id", employeeController.Get)
	router.Post("/employees", employeeController.Create)
	router.Patch("/employees/:id", employeeController.Update)
	router.Delete("/employees/:id", employeeController.Delete)
}
