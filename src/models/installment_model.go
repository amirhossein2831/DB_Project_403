package models

import "time"

type Installment struct {
	ID           int        `json:"id" sql:"id"`
	LoanID       int        `json:"loan_id" sql:"loan_id"`
	AmountPaid   float64    `json:"amount_paid" sql:"amount_paid"`
	InterestPaid float64    `json:"interest_paid" sql:"interest_paid"`
	TotalPaid    float64    `json:"total_paid" sql:"total_paid"`
	DueDate      time.Time  `json:"due_date" sql:"due_date"`
	PaidDate     *time.Time `json:"paid_date" sql:"paid_date"`
	Loan         *Loan      `json:"loan" sql:""`
}
