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

func (repository *LoanRepository) List(status string) ([]*models.Loan, error) {
	query := "SELECT * FROM loan l INNER JOIN customer c ON l.customer_id = c.id"

	var args []interface{}
	if status != "" && (status == string(models.PendingLoanStatus) || status == string(models.ApprovedLoanStatus) || status == string(models.RepaidLoanStatus) || status == string(models.DefaultedLoanStatus)) || status == string(models.FinishedLoanStatus) {
		query += " WHERE l.status = $1"
		args = append(args, status)
	}

	var loans []*models.Loan
	rows, err := pgx.GetInstance().Query(context.Background(), query, args...)
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

func (repository *LoanRepository) ListWithMinInstallmentsPaid() ([]*models.LoanWithMinInstallmentPain, error) {
	var customers []*models.LoanWithMinInstallmentPain
	rows, err := pgx.GetInstance().Query(context.Background(), `
	WITH TotalInstallments AS (
		SELECT
			COUNT(*) AS insallment_number,
			loan_id
		FROM
			installment
		WHERE
			paid_date IS NOT NULL
		GROUP BY
			loan_id
	),
		 MinInstallments AS (
			 SELECT
				 MIN(insallment_number) AS min_insallment_number
			 FROM
				 TotalInstallments
		 )
	SELECT
		t.loan_id as id,
		l.customer_id,
		l.type,
		l.status,
		l.amount,
		t.insallment_number as total_installments_paid
	FROM
		TotalInstallments t
			CROSS JOIN MinInstallments m
			INNER JOIN loan l ON t.loan_id = l.id
	WHERE
		t.insallment_number = m.min_insallment_number;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.LoanWithMinInstallmentPain
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *LoanRepository) Get(id string) (*models.Loan, error) {
	var loan models.Loan
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT * FROM loan l LEFT JOIN customer c ON l.customer_id = c.id WHERE l.id = $1", id)
	err := utils.FillStructFromRowWithJoin(row, &loan)
	if err != nil {
		return nil, err
	}

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
