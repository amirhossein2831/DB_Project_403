package services

import (
	"DB_Project/src/models"
	"DB_Project/src/repositories"
)

type ViewService struct {
	Repository *repositories.ViewRepository
}

func NewViewService() *ViewService {
	return &ViewService{
		Repository: repositories.NewViewRepository(),
	}
}

func (service *ViewService) GetCustomerAccountsView() ([]*models.CustomerAccounts, error) {
	return service.Repository.CustomerAccountsView()
}

func (service *ViewService) GetBankTransactionView() ([]*models.BankTransactions, error) {
	return service.Repository.BankTransactionsView()
}

func (service *ViewService) GetBankMemberView() ([]*models.BankMembers, error) {
	return service.Repository.BankMemberView()
}
