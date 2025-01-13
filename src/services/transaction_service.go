package services

import (
	"DB_Project/src/api/http/requests/transaction"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"context"
)

type TransactionService struct {
	Repository *repositories.TransactionRepository
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		Repository: repositories.NewTransactionRepository(),
	}
}

func (service *TransactionService) GetTransactions(ctx context.Context) ([]*models.Transaction, error) {
	sourceId := ctx.Value("source_id").(int)

	return service.Repository.List(sourceId)
}

func (service *TransactionService) GetTransaction(id string) (*models.Transaction, error) {
	return service.Repository.Get(id)
}

func (service *TransactionService) CreateTransaction(req *transaction.CreateTransactionRequest) error {
	// profile
	transaction := &models.Transaction{
		Type:                 req.Type,
		Amount:               req.Amount,
		SourceAccountId:      req.SourceAccountId,
		DestinationAccountId: req.DestinationAccountId,
	}

	return service.Repository.Create(transaction)
}

func (service *TransactionService) UpdateTransaction(req *transaction.UpdateTransactionRequest, id string) error {
	if req.Type != nil && *req.Type != "" {
		err := service.Repository.UpdateField("type", id, *req.Type)
		if err != nil {
			return err
		}
	}

	if req.Amount != nil {
		err := service.Repository.UpdateField("amount", id, *req.Amount)
		if err != nil {
			return err
		}
	}

	if req.SourceAccountId != nil {
		err := service.Repository.UpdateField("source_account_id", id, *req.SourceAccountId)
		if err != nil {
			return err
		}
	}

	err := service.Repository.UpdateField("destination_account_id", id, req.DestinationAccountId)
	if err != nil {
		return err
	}

	return nil
}

func (service *TransactionService) DeleteTransaction(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
