package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func TransactionSeeder() {
	log.Println("Transaction Customer table...")

	query := `
	INSERT INTO transaction (id, type, amount, source_account_id, destination_account_id, created_at) VALUES
		(1, 'deposit', 1000.00, NULL, 1, '2024-01-01 10:00:00'),
		(2, 'withdrawal', 500.00, 1, NULL, '2024-01-02 10:00:00'),
		(3, 'transfer', 200.00, 1, 2, '2024-01-03 10:00:00'),
		(4, 'transfer', 300.00, 2, 3, '2024-01-04 10:00:00'),
		(5, 'deposit', 1500.00, NULL, 4, '2024-01-05 10:00:00');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Transaction seeded successfully.")
}
