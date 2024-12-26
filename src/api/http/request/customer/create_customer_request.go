package customer

import (
	"DB_Project/src/models"
)

type CreateCustomerRequest struct {
	FirstName    string              `json:"first_name" validate:"required"`
	LastName     string              `json:"last_name" validate:"required"`
	BirthDate    string              `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Phone        string              `json:"phone" validate:"required"`
	Email        string              `json:"email" validate:"required,email"`
	Address      string              `json:"address" validate:"required"`
	CustomerType models.CustomerType `json:"customer_type" validate:"required,oneof=individual legal_entity"`
}
