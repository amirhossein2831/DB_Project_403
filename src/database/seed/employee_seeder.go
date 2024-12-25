package seed

import (
	"DB_Project/src/database/connection/pgx"
	"context"
	"log"
)

func EmployeeSeeder() {
	log.Println("Employee Customer table...")

	query := `
	INSERT INTO employee (profile_id, position)VALUES
		(1, 'Manager'),
		(2, 'Cashier'),
		(3, 'Loan Officer'),
		(4, 'IT Specialist'),
		(5, 'Accountant');
	`

	// Execute the query
	_, err := pgx.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Employee seeded successfully.")
}
