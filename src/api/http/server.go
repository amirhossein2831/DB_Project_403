package http

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"log"
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
		app = fiber.New()

		// append middleware here
		app.Use(recover.New())
		app.Use(requestid.New())
	})
}

func initServer() error {
	// Initial v1 routes
	getNewRouter()

	log.Fatal(app.Listen(
		fmt.Sprintf("%s:%s", os.Getenv("APP_NAME"), os.Getenv("APP_PORT")),
	))

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
