package services

import (
	"DB_Project/src/api/http/request/customer"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"time"
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

func (service *CustomerService) CreateCustomer(req *customer.CreateCustomerRequest) error {
	// cast birth_date to date
	birthDate, _ := time.Parse("2006-01-02", req.BirthDate)

	// customer
	customer := &models.Customer{
		Type: req.CustomerType,
	}

	// profile
	profile := &models.Profile{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: birthDate,
		Phone:     req.Phone,
		Email:     req.Email,
		Address:   req.Address,
	}

	return service.Repository.Create(customer, profile)
}

func (service *CustomerService) DeleteCustomer(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
