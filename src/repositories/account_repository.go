package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (repository *AccountRepository) List() ([]*models.Account, error) {
	var accounts []*models.Account
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT *  FROM account a INNER JOIN customer c ON a.customer_id =  c.id ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err = utils.FillStructFromRowsWithJoin(rows, &account)
		accounts = append(accounts, &account)
	}

	return accounts, rows.Err()
}

func (repository *AccountRepository) ListByCustomerId(customerId int) ([]*models.Account, error) {
	var accounts []*models.Account
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT *  FROM account WHERE customer_id=$1", customerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.ID, &account.AccountNumber, &account.Type, &account.Amount, &account.Status, &account.CustomerID, &account.CreatedAt, &account.ClosedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return accounts, rows.Err()
}

func (repository *AccountRepository) Get(id string) (*models.Account, error) {
	var account models.Account
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT * FROM account a LEFT JOIN customer c ON a.customer_id=c.id WHERE a.id =$1", id)
	err := utils.FillStructFromRowWithJoin(row, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) Create(account *models.Account) error {
	// Insert account and retrieve the generated account ID
	_, err := pgx.GetInstance().Exec(context.Background(),
		"INSERT INTO account (account_number, type, amount, status, customer_id) VALUES ($1, $2, $3, $4, $5)",
		account.AccountNumber, account.Type, account.Amount, account.Status, account.CustomerID,
	)
	return err
}

func (repository *AccountRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE account SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *AccountRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM account WHERE id=$1", id)
	return err
}
