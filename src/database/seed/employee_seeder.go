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
		(6, 'Manager'),
		(7, 'Cashier'),
		(8, 'Loan Officer'),
		(9, 'IT Specialist'),
		(10, 'Accountant');
	`

	// Execute the query
	_, err := pgx.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Employee seeded successfully.")
}
