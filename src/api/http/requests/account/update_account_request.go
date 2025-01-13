package account

import "DB_Project/src/models"

type UpdateAccountRequest struct {
	AccountNumber *string               `json:"account_number" validate:"omitempty"`
	Type          *models.AccountType   `json:"type" validate:"omitempty,oneof='' 'savings' 'current' 'business'"`
	Amount        *float64              `json:"amount" validate:"omitempty,gt=0"`
	Status        *models.AccountStatus `json:"status" validate:"omitempty,oneof='' 'active' 'closed' 'pending"`
	CustomerID    *int                  `json:"customer_id" validate:"omitempty,gt=0"`
}
