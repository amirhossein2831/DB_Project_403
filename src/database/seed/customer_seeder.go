package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func CustomerSeeder() {
	log.Println("seeding Customer table...")

	query := `
	INSERT INTO customer (profile_id, type)	VALUES
		(1, 'individual'),
		(2, 'legal_entity'),
		(3, 'individual'),
		(4, 'legal_entity'),
		(5, 'individual');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Customer seeded successfully.")
}
