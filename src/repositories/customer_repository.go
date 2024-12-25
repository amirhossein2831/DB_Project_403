package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type CustomerRepository struct {
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (repository *CustomerRepository) List() ([]*models.Customer, error) {
	var customers []*models.Customer
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT c.*, p.*, a.* FROM customer c LEFT JOIN profile p ON c.profile_id = p.id RIGHT JOIN account a ON c.id = a.customer_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer
		err = utils.FillStructFromRowsWithJoin(rows, &customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	return models.FetchCustomersAccount(customers), rows.Err()
}

func (repository *CustomerRepository) Get(id string) (*models.Customer, error) {
	var customers []*models.Customer
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT c.*, p.*, a.* FROM customer c LEFT JOIN profile p ON c.profile_id = p.id RIGHT JOIN account a ON c.id = a.customer_id WHERE c.id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var customer models.Customer
		err = utils.FillStructFromRowsWithJoin(rows, &customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	return models.FetchCustomerAccount(customers)
}

func (repository *CustomerRepository) Create(customer *models.Customer, profile *models.Profile) error {
	// context
	ctx := context.Background()

	// Start a transaction
	tx, err := pgx.GetInstance().Begin(ctx)
	if err != nil {
		return err
	}
	// Ensure that the transaction is rolled back if any error occurs
	defer tx.Rollback(ctx)

	// Insert profile and retrieve the generated profile ID
	err = tx.QueryRow(ctx, "INSERT INTO profile (first_name, last_name, birth_date, phone, email, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		profile.FirstName, profile.LastName, profile.BirthDate, profile.Phone, profile.Email, profile.Address,
	).Scan(&profile.ID)
	if err != nil {
		return err
	}

	// Insert customer using the generated profile ID
	_, err = tx.Exec(ctx, "INSERT INTO customer (profile_id, type) VALUES ($1, $2)",
		profile.ID, customer.Type,
	)
	if err != nil {
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit(ctx)
}

func (repository *CustomerRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE customer SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *CustomerRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM customer WHERE id=$1", id)
	return err
}
