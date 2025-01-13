package services

import (
	"DB_Project/src/api/http/requests/customer"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"time"
)

type CustomerService struct {
	Repository        *repositories.CustomerRepository
	ProfileRepository *repositories.ProfileRepository
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		Repository:        repositories.NewCustomerRepository(),
		ProfileRepository: repositories.NewProfileRepository(),
	}
}

func (service *CustomerService) GetCustomers() ([]*models.Customer, error) {
	return service.Repository.List()
}

func (service *CustomerService) GetCustomersWithFullName() ([]*models.CustomerWithFullName, error) {
	return service.Repository.ListWithFullName()
}

func (service *CustomerService) GetCustomersWithFullNameAndTotalAmount() ([]*models.CustomerWithFullNameAndAmount, error) {
	return service.Repository.ListWithFullNameAndTotalAmount()
}

func (service *CustomerService) GetCustomerWithFullNameAndAccountNumber() ([]*models.CustomerWithAccountNumber, error) {
	return service.Repository.ListWithFullNameAndAccountNumber()
}

func (service *CustomerService) GetCustomerWithMostLoan() ([]*models.CustomerWithMostLoan, error) {
	return service.Repository.ListWithMostLoan()
}

func (service *CustomerService) GetCustomerWithInstallmentsPenalty() ([]*models.CustomerWithPenaltyInstallment, error) {
	return service.Repository.ListWithInstallmentsPenalty()
}

func (service *CustomerService) GetCustomerWithMostAmount() ([]*models.CustomerWithMostAmount, error) {
	return service.Repository.ListWithMostAmount()
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

// TODO: use one query to update several field
func (service *CustomerService) UpdateCustomer(req *customer.UpdateCustomerRequest, id string) error {
	if req.FirstName != nil && *req.FirstName != "" {
		err := service.ProfileRepository.UpdateField("first_name", id, *req.FirstName)
		if err != nil {
			return err
		}
	}

	if req.LastName != nil && *req.LastName != "" {
		err := service.ProfileRepository.UpdateField("last_name", id, *req.LastName)
		if err != nil {
			return err
		}
	}

	if req.BirthDate != nil && *req.BirthDate != "" {
		err := service.ProfileRepository.UpdateField("birth_date", id, *req.BirthDate)
		if err != nil {
			return err
		}
	}

	if req.Phone != nil && *req.Phone != "" {
		err := service.ProfileRepository.UpdateField("phone", id, *req.Phone)
		if err != nil {
			return err
		}
	}

	if req.Email != nil && *req.Email != "" {
		err := service.ProfileRepository.UpdateField("email", id, *req.Email)
		if err != nil {
			return err
		}
	}

	if req.Address != nil && *req.Address != "" {
		err := service.ProfileRepository.UpdateField("address", id, *req.Address)
		if err != nil {
			return err
		}
	}

	if req.CustomerType != nil && *req.CustomerType != "" {
		err := service.Repository.UpdateField("type", id, *req.CustomerType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *CustomerService) DeleteCustomer(id string) error {
	res, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.ProfileRepository.Delete(res.ProfileID)
}
