package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func InstallmentSeeder() {
	log.Println("Installment Customer table...")

	query := `
	INSERT INTO installment (loan_id, amount_paid, interest_paid, total_paid, due_date, paid_date)VALUES
		(1, 100.00, 10.00, 110.00, '2024-02-01', '2024-02-01'),
		(1, 100.00, 10.00, 110.00, '2024-03-01', '2024-03-01'),
		(2, 500.00, 100.00, 600.00, '2024-04-01', '2024-04-01'),
		(2, 500.00, 100.00, 600.00, '2024-05-01', '2024-05-01'),
		(3, 250.00, 50.00, 300.00, '2024-06-01', '2024-06-01');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Installment seeded successfully.")
}
