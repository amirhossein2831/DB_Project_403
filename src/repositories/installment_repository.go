package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type InstallmentRepository struct {
}

func NewInstallmentRepository() *InstallmentRepository {
	return &InstallmentRepository{}
}

func (repository *InstallmentRepository) List() ([]*models.Installment, error) {
	var installments []*models.Installment
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM installment i INNER JOIN loan l ON i.loan_id = l.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var installment models.Installment
		err = utils.FillStructFromRowsWithJoin(rows, &installment)
		installments = append(installments, &installment)
	}

	return installments, rows.Err()
}

func (repository *InstallmentRepository) Get(id string) (*models.Installment, error) {
	var installment models.Installment
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT * FROM installment i LEFT JOIN loan l ON i.loan_id = l.id WHERE i.id = $1", id)
	err := utils.FillStructFromRowWithJoin(row, &installment)
	if err != nil {
		return nil, err
	}

	return &installment, nil
}

func (repository *InstallmentRepository) Create(installment *models.Installment) error {
	// Insert installment and retrieve the generated installment ID
	_, err := pgx.GetInstance().Exec(context.Background(),
		"INSERT INTO installment (loan_id, amount_paid, interest_paid, total_paid, due_date, paid_date) VALUES ($1, $2, $3, $4, $5, $6)",
		installment.LoanID, installment.AmountPaid, installment.InterestPaid, installment.TotalPaid, installment.DueDate, installment.PaidDate,
	)
	return err
}

func (repository *InstallmentRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE installment SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *InstallmentRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM installment WHERE id = $1", id)
	return err
}
