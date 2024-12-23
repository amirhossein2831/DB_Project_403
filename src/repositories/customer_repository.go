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

func (repository *CustomerRepository) List() ([]*models.Customer, error) {
	var customers []*models.Customer
	rows, err := database.GetInstance().Query(context.Background(), "SELECT c.*, p.* FROM customer c LEFT JOIN profile p ON c.profile_id = p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer
		err = utils.FillStructFromRowsWithJoin(rows, &customer)
		customers = append(customers, &customer)
	}

	return customers, rows.Err()
}

func (repository *CustomerRepository) Get(id string) (*models.Customer, error) {
	var customer models.Customer
	row := database.GetInstance().QueryRow(context.Background(), "SELECT c.*, p.* FROM customer c LEFT JOIN profile p ON c.profile_id = p.id WHERE c.id=$1", id)
	err := utils.FillStructFromRowWithJoin(row, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (repository *CustomerRepository) Delete(id string) error {
	_, err := database.GetInstance().Exec(context.Background(), "DELETE FROM customer WHERE id=$1", id)
	return err
}
