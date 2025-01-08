package services

import (
	customer "DB_Project/src/api/http/request/employee"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"time"
)

type EmployeeService struct {
	Repository        *repositories.EmployeeRepository
	ProfileRepository *repositories.ProfileRepository
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		Repository:        repositories.NewEmployeeRepository(),
		ProfileRepository: repositories.NewProfileRepository(),
	}
}
func (service *EmployeeService) GetEmployees() ([]*models.Employee, error) {
	return service.Repository.List()
}
func (service *EmployeeService) GetEmployee(id string) (*models.Employee, error) {
	return service.Repository.Get(id)
}
func (service *EmployeeService) CreateEmployee(req *customer.CreateEmployeeRequest) error {
	// cast birth_date to date
	birthDate, _ := time.Parse("2006-01-02", req.BirthDate)

	// employee
	employee := &models.Employee{
		Position: req.Position,
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

	return service.Repository.Create(employee, profile)
}

// TODO: use one query to update several field
func (service *EmployeeService) UpdateEmployee(req *customer.UpdateEmployeeRequest, id string) error {
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

	if req.Position != nil && *req.Position != "" {
		err := service.Repository.UpdateField("position", id, *req.Position)
		if err != nil {
			return err
		}
	}

	return nil
}
func (service *EmployeeService) DeleteEmployee(id string) error {
	res, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.ProfileRepository.Delete(res.ProfileID)
}
