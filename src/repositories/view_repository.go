package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
)

type ViewRepository struct {
}

func NewViewRepository() *ViewRepository {
	return &ViewRepository{}
}

func (repository *ViewRepository) CustomerAccountsView() ([]*models.CustomerAccounts, error) {
	var customers []*models.CustomerAccounts
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM customer_accounts")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerAccounts
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *ViewRepository) BankTransactionsView() ([]*models.BankTransactions, error) {
	var customers []*models.BankTransactions
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM bank_transactions")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.BankTransactions
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *ViewRepository) BankMemberView() ([]*models.BankMembers, error) {
	var customers []*models.BankMembers
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM bank_members")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.BankMembers
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}
