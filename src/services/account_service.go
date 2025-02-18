package services

import (
	"DB_Project/src/api/http/requests/account"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"context"
)

type AccountService struct {
	Repository *repositories.AccountRepository
}

func NewAccountService() *AccountService {
	return &AccountService{
		Repository: repositories.NewAccountRepository(),
	}
}

func (service *AccountService) GetAccounts(ctx context.Context) ([]*models.Account, error) {
	status := ctx.Value("status").(string)
	minAmount := ctx.Value("min_amount").(float64)

	return service.Repository.List(status, minAmount)
}

func (service *AccountService) GetAccount(id string) (*models.Account, error) {
	return service.Repository.Get(id)
}

func (service *AccountService) CreateAccount(req *account.CreateAccountRequest) error {
	// profile
	account := &models.Account{
		AccountNumber: req.AccountNumber,
		Type:          req.Type,
		Amount:        req.Amount,
		Status:        req.Status,
		CustomerID:    req.CustomerID,
	}

	return service.Repository.Create(account)
}

func (service *AccountService) UpdateAccount(req *account.UpdateAccountRequest, id string) error {
	if req.AccountNumber != nil && *req.AccountNumber != "" {
		err := service.Repository.UpdateField("account_number", id, *req.AccountNumber)
		if err != nil {
			return err
		}
	}

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

	if req.Status != nil && *req.Status != "" {
		err := service.Repository.UpdateField("status", id, *req.Status)
		if err != nil {
			return err
		}
	}

	if req.CustomerID != nil {
		err := service.Repository.UpdateField("customer_id", id, *req.CustomerID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *AccountService) DeleteAccount(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
