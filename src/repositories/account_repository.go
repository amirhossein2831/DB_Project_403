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
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT c.* ,a.*  FROM account a LEFT JOIN customer c ON a.customer_id =  c.id = a.customer_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err = utils.FillStructFromRows(rows, &account)
		accounts = append(accounts, &account)
	}

	return accounts, rows.Err()
}

func (repository *AccountRepository) Get(id string) (*models.Account, error) {
	var account models.Account
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT c.* ,a.*  FROM account a LEFT JOIN customer c ON a.customer_id =  c.id = a.customer_id WHERE a.id=$1", id)
	err := utils.FillStructFromRow(row, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repository *AccountRepository) Create(account *models.Account) error {
	// Insert account and retrieve the generated account ID
	return pgx.GetInstance().QueryRow(context.Background(),
		"INSERT INTO account (account_number, type, amount, status, customer_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		account.AccountNumber, account.Type, account.Amount, account.Status, account.CustomerID,
	).Scan(&account.ID)
}

func (repository *AccountRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE account SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *AccountRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM account WHERE id=$1", id)
	return err
}
