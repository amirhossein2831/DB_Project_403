package account

import (
	"DB_Project/src/models"
)

type CreateAccountRequest struct {
	AccountNumber string               `json:"account_number" validate:"required"`
	Type          models.AccountType   `json:"type" validate:"required,oneof=savings current business"`
	Amount        float64              `json:"amount" validate:"required,gt=0"`
	Status        models.AccountStatus `json:"status" validate:"required,oneof=active closed pending"`
	CustomerID    int                  `json:"customer_id" validate:"required,gt=0"`
}
