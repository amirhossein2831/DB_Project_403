package migrate

import (
	"DB_Project/src/database"
	"DB_Project/src/database/migrations"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// Migrate Commands for interacting with database
var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Commands for migrate tables",
}

func init() {
	Migrate.AddCommand(
		migrateUp,
		migrateDown,
	)
}

var migrateUp = &cobra.Command{
	Use:   "up",
	Short: "migrate the tables up",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		log.Println("Starting migrate up...    " + "timestamp" + time.Now().String())
		err := migrations.MigrateUp()
		if err != nil {
			log.Fatalf("Failed to migrate up: %v.    timestamp: %s", err, time.Now().String())
		}
		log.Println("migrate up successfully.    " + "timestamp" + time.Now().String())

		afterMigrate()
	},
}

var migrateDown = &cobra.Command{
	Use:   "down",
	Short: "migrate the tables down",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		log.Println("Starting migrate down...    " + "timestamp" + time.Now().String())
		err := migrations.MigrateDown()
		if err != nil {
			log.Fatalf("Failed to migrate down: %v.    timestamp: %s", err, time.Now().String())
		}
		log.Println("migrate down successfully.    " + "timestamp" + time.Now().String())

		afterMigrate()
	},
}

func beforeMigrate() {
	// Initialize Env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ENV Service: Failed to  loading .env file. %v.    timestamp: %s", err, time.Now().String())
	}

	// Initialize Database
	err = database.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize: %v.    timestamp: %s", err, time.Now().String())
	}
}

func afterMigrate() {
	err := database.Close()
	if err != nil {
		log.Fatalf("Databasde Service: Failed to close database. %v.    timestamp: %s \n", err, time.Now().String())
	}
}
