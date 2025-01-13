package make

import (
	"DB_Project/cmd/make/all"
	"DB_Project/cmd/make/controller"
	"DB_Project/cmd/make/exception"
	"DB_Project/cmd/make/repository"
	"DB_Project/cmd/make/service"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// Make Commands for making file
var Make = &cobra.Command{
	Use:   "make",
	Short: "Commands for making file",
}

func init() {
	// init requirement
	beforeMake()

	// init subcommand
	Make.AddCommand(
		exception.ExceptionCmd,
		controller.ControllerCmd,
		service.ServiceCmd,
		repository.RepositoryCmd,
		all.AllCmd,
	)
}

func beforeMake() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
	}
}
