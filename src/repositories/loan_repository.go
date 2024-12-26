package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type LoanRepository struct {
	InstallmentRepository *InstallmentRepository
}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{
		InstallmentRepository: NewInstallmentRepository(),
	}
}

func (repository *LoanRepository) List() ([]*models.Loan, error) {
	var loans []*models.Loan
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM loan l INNER JOIN customer c ON l.customer_id = c.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var loan models.Loan
		err = utils.FillStructFromRowsWithJoin(rows, &loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	for _, loan := range loans {
		var installments []*models.Installment
		installments, err = repository.InstallmentRepository.ListOfInstallmentByLoanId(loan.ID)
		if err != nil {
			return nil, err
		}
		loan.Installments = installments
	}

	return loans, rows.Err()
}

func (repository *LoanRepository) Get(id string) (*models.Loan, error) {
	var loan models.Loan
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT * FROM loan l LEFT JOIN customer c ON l.customer_id = c.id WHERE l.id = $1", id)
	err := utils.FillStructFromRowWithJoin(row, &loan)

	var installments []*models.Installment
	installments, err = repository.InstallmentRepository.ListOfInstallmentByLoanId(loan.ID)
	if err != nil {
		return nil, err
	}
	loan.Installments = installments

	return &loan, err
}

func (repository *LoanRepository) Create(loan *models.Loan) error {
	// Insert loan record
	_, err := pgx.GetInstance().Exec(context.Background(),
		`INSERT INTO loan (customer_id, type, status, amount, interest_rate, repayment_period, finished_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		loan.CustomerID, loan.Type, loan.Status, loan.Amount, loan.InterestRate, loan.RepaymentPeriod, loan.FinishedAt,
	)
	return err
}

func (repository *LoanRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE loan SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *LoanRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM loan WHERE id=$1", id)
	return err
}
