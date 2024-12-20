package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func LoanSeeder() {
	log.Println("Loan Customer table...")

	query := `
	INSERT INTO loan (id, customer_id, type, status, amount, interest_rate, repayment_period, created_at, finished_at)VALUES
		(1, 1, 'personal', 'active', 2000.00, 5.5, 12, '2024-01-10 10:00:00', NULL),
		(2, 2, 'mortgage', 'active', 100000.00, 3.0, 240, '2024-02-15 10:00:00', NULL),
		(3, 3, 'auto', 'inactive', 15000.00, 4.2, 60, '2024-03-20 10:00:00', '2024-09-20 10:00:00'),
		(4, 4, 'business', 'active', 50000.00, 6.0, 36, '2024-04-25 10:00:00', NULL),
		(5, 5, 'personal', 'inactive', 3000.00, 5.0, 12, '2024-05-30 10:00:00', '2024-11-30 10:00:00');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Loan seeded successfully.")
}
