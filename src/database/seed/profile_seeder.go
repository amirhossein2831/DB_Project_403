package seed

import (
	"DB_Project/src/database/connection/pgx"
	"context"
	"log"
)

func ProfileSeeder() {
	log.Println("seeding Profile table...")

	query := `
	INSERT INTO profile (first_name, last_name, birth_date, phone, email, address) VALUES
		('John', 'Doe', '1985-07-15', '1234567890', 'john.doe@example.com', '123 Main St'),
		('Jane', 'Smith', '1990-03-22', '0987654321', 'jane.smith@example.com', '456 Elm St'),
		('Emily', 'Brown', '1987-11-11', '1112223333', 'emily.brown@example.com', '789 Oak St'),
		('Michael', 'Johnson', '1992-05-30', '4445556666', 'michael.johnson@example.com', '321 Pine St'),
		('Sarah', 'Williams', '1995-09-10', '7778889999', 'sarah.williams@example.com', '654 Cedar St');
	`

	// Execute the query
	_, err := pgx.GetInstance().Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Error seeding profiles: %v", err)
	}

	log.Println("Profiles seeded successfully.")
}
