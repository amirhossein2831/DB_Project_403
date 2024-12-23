package repositories

import (
	"DB_Project/src/database"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
)

type CustomerRepository struct {
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (repository *CustomerRepository) GetCustomer() ([]*models.Customer, error) {
	var customers []*models.Customer
	db := database.GetInstance()
	rows, err := db.Query(context.Background(), "SELECT * FROM customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer
		err = utils.FillStructFromRow(rows, &customer)
		customers = append(customers, &customer)
	}

	return customers, rows.Err()
}
