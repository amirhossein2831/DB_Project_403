package seed

import (
	"DB_Project/src/database"
	"DB_Project/src/database/seed"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// Seed Commands for interacting with database
var Seed = &cobra.Command{
	Use:   "seed",
	Short: "Commands for seed tables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		beforeMigrate()

		// add seeders
		log.Println("Starting seeding table...    " + "timestamp" + time.Now().String())
		seed.ProfileSeeder()
		seed.CustomerSeeder()
		seed.EmployeeSeeder()
		seed.AccountSeeder()
		seed.TransactionSeeder()
		seed.LoanSeeder()
		seed.InstallmentSeeder()

		log.Println("seeding table successfully.    " + "timestamp" + time.Now().String())

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
