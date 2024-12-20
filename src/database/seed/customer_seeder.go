package seed

import (
	"DB_Project/src/database"
	"context"
	"log"
)

func CustomerSeeder() {
	log.Println("seeding Customer table...")

	query := `
	INSERT INTO customer (id, profile_id, type)	VALUES
		(1, 1, 'regular'),
		(2, 2, 'premium'),
		(3, 3, 'regular'),
		(4, 4, 'corporate'),
		(5, 5, 'premium');
	`

	// Execute the query
	_, err := database.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Customer seeded successfully.")
}
