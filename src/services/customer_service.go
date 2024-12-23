package services

import (
	"DB_Project/src/models"
	"DB_Project/src/repositories"
)

type CustomerService struct {
	Repository *repositories.CustomerRepository
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		Repository: repositories.NewCustomerRepository(),
	}
}

func (service *CustomerService) GetCustomers() ([]*models.Customer, error) {
	return service.Repository.GetCustomers()
}

func (service *CustomerService) GetCustomer(id string) (*models.Customer, error) {
	return service.Repository.GetCustomer(id)
}
