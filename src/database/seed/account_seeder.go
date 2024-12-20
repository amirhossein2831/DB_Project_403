package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func AccountSeeder() {
	log.Println("Account Customer table...")

	query := `
	INSERT INTO account (id, account_number, type, amount, status, customer_id, created_at, closed_at) VALUES
		(1, 'ACC10001', 'savings', 5000.00, 'active', 1, '2024-01-01 10:00:00', NULL),
		(2, 'ACC10002', 'current', 10000.00, 'active', 2, '2024-02-01 10:00:00', NULL),
		(3, 'ACC10003', 'savings', 3000.00, 'active', 3, '2024-03-01 10:00:00', NULL),
		(4, 'ACC10004', 'current', 8000.00, 'active', 4, '2024-04-01 10:00:00', NULL),
		(5, 'ACC10005', 'savings', 7000.00, 'inactive', 5, '2024-05-01 10:00:00', '2024-06-01 10:00:00');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Account seeded successfully.")
}
