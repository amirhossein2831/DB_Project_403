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
	return service.Repository.List()
}

func (service *CustomerService) GetCustomer(id string) (*models.Customer, error) {
	return service.Repository.Get(id)
}

func (service *CustomerService) DeleteCustomer(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
