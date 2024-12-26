package loan

import (
	"DB_Project/src/models"
	"time"
)

type UpdateLoanRequest struct {
	CustomerID      *int               `json:"customer_id" validate:"omitempty,gt=0"`
	Type            *models.LoanType   `json:"type" validate:"omitempty,oneof='' 'personal' 'mortgage' 'business'"`
	Status          *models.LoanStatus `json:"status" validate:"omitempty,oneof='' 'pending' 'approved' 'repaid' 'defaulted'"`
	Amount          *float64           `json:"amount" validate:"omitempty,gt=0"`
	InterestRate    *float32           `json:"interest_rate" validate:"omitempty,gt=0"`
	RepaymentPeriod *int               `json:"repayment_period" validate:"omitempty,gt=0"`
	FinishedAt      *time.Time         `json:"finished_at,omitempty" validate:"omitempty"`
}
