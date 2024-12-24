package repositories

import (
	"DB_Project/src/database"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type EmployeeRepository struct {
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{}
}

func (repository *EmployeeRepository) List() ([]*models.Employee, error) {
	var employees []*models.Employee
	rows, err := database.GetInstance().Query(context.Background(), "SELECT e.*, p.* FROM employee e LEFT JOIN profile p ON e.profile_id = p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err = utils.FillStructFromRowsWithJoin(rows, &employee)
		employees = append(employees, &employee)
	}

	return employees, rows.Err()
}

func (repository *EmployeeRepository) Get(id string) (*models.Employee, error) {
	var employee models.Employee
	row := database.GetInstance().QueryRow(context.Background(), "SELECT e.*, p.* FROM employee e LEFT JOIN profile p ON e.profile_id = p.id WHERE e.id=$1", id)
	err := utils.FillStructFromRowWithJoin(row, &employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (repository *EmployeeRepository) Create(employee *models.Employee, profile *models.Profile) error {
	// context
	ctx := context.Background()

	// Start a transaction
	tx, err := database.GetInstance().Begin(ctx)
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

	// Insert employee using the generated profile ID
	_, err = tx.Exec(ctx, "INSERT INTO employee (profile_id, position) VALUES ($1, $2)",
		profile.ID, employee.Position,
	)
	if err != nil {
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit(ctx)
}

func (repository *EmployeeRepository) UpdateField(name, id string, value interface{}) error {
	_, err := database.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE employee SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *EmployeeRepository) Delete(id string) error {
	_, err := database.GetInstance().Exec(context.Background(), "DELETE FROM employee WHERE id=$1", id)
	return err
}
