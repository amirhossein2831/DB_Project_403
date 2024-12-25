package http

import (
	"DB_Project/src/api/http/routes"
	"DB_Project/src/pkg/validation"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"os"
	"sync"
)

var app *fiber.App
var once sync.Once

func Init() error {
	return initServer()
}

func getNewRouter() {
	once.Do(func() {
		app = fiber.New(fiber.Config{
			StructValidator: &validation.StructValidator{Validator: validator.New()},
		})

		// append middleware here
		app.Use(recover.New())
		app.Use(requestid.New())
		app.Use(logger.New())
	})
}

func initServer() error {
	getNewRouter()

	// Initial v1 routes
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			routes.CustomerRoute(v1)
			routes.EmployeeRoute(v1)
			routes.AccountRoute(v1)
		}
	}

	err := app.Listen(
		fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")),
	)
	if err != nil {
		return err
	}

	return nil
}

func ShutdownServer() error {
	if app == nil {
		return nil
	}

	err := app.Shutdown()
	if err != nil {
		return err
	}

	return nil
}
