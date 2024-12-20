package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func EmployeeSeeder() {
	log.Println("Employee Customer table...")

	query := `
	INSERT INTO employee (id, profile_id, position)VALUES
		(1, 1, 'Manager'),
		(2, 2, 'Cashier'),
		(3, 3, 'Loan Officer'),
		(4, 4, 'IT Specialist'),
		(5, 5, 'Accountant');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Employee seeded successfully.")
}
