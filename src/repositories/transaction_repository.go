package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type TransactionRepository struct {
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (repository *TransactionRepository) List() ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT *  FROM transaction t  INNER JOIN account sa ON t.source_account_id = sa.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err = utils.FillStructFromRowsWithJoin(rows, &transaction)
		transactions = append(transactions, &transaction)
	}

	for _, transaction := range transactions {
		if transaction.DestinationAccountId != nil {
			var account models.Account
			destinationAccountsRow := pgx.GetInstance().QueryRow(context.Background(), "SELECT * from account WHERE id = $1", transaction.DestinationAccountId)
			err = utils.FillStructFromRow(destinationAccountsRow, &account)
			if err != nil {
				return nil, err
			}
			transaction.DestinationAccount = &account
		}

	}

	return transactions, rows.Err()
}

func (repository *TransactionRepository) Get(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT *  FROM transaction t  INNER JOIN account sa ON t.source_account_id = sa.id WHERE t.id=$1", id)
	err := utils.FillStructFromRowWithJoin(row, &transaction)
	if err != nil {
		return nil, err
	}

	if transaction.DestinationAccountId != nil {
		var account models.Account
		destinationAccountsRow := pgx.GetInstance().QueryRow(context.Background(), "SELECT * from account WHERE id = $1", transaction.DestinationAccountId)
		err = utils.FillStructFromRow(destinationAccountsRow, &account)
		if err != nil {
			return nil, err
		}
		transaction.DestinationAccount = &account
	}

	return &transaction, nil
}

func (repository *TransactionRepository) Create(transaction *models.Transaction) error {
	// Insert account and retrieve the generated account ID
	_, err := pgx.GetInstance().Exec(context.Background(),
		"INSERT INTO transaction (type, amount, source_account_id, destination_account_id) VALUES ($1, $2, $3, $4)",
		transaction.Type, transaction.Amount, transaction.SourceAccountId, transaction.DestinationAccountId,
	)
	return err
}

func (repository *TransactionRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE transaction SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *TransactionRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM transaction WHERE id=$1", id)
	return err
}
