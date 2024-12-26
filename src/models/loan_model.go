package models

import (
	"time"
)

type LoanType string
type LoanStatus string

const (
	PersonalLoanType LoanType = "personal"
	MortgageLoanType LoanType = "mortgage"
	BusinessLoanType LoanType = "business"
)

const (
	PendingLoanStatus   LoanStatus = "pending"
	ApprovedLoanStatus  LoanStatus = "approved"
	RepaidLoanStatus    LoanStatus = "repaid"
	DefaultedLoanStatus LoanStatus = "defaulted"
)

type Loan struct {
	ID              int            `json:"id" sql:"id"`
	CustomerID      int            `json:"customer_id" sql:"customer_id"`
	Type            LoanType       `json:"type" sql:"type"`
	Status          LoanStatus     `json:"status" sql:"status"`
	Amount          float64        `json:"amount" sql:"amount"`
	InterestRate    float32        `json:"interest_rate" sql:"interest_rate"`
	RepaymentPeriod int            `json:"repayment_period" sql:"repayment_period"`
	CreatedAt       time.Time      `json:"created_at" sql:"created_at"`
	FinishedAt      time.Time      `json:"finished_at" sql:"finished_at"`
	Customer        *Customer      `json:"customer" sql:""`
	Installments    []*Installment `json:"installments" sql:"-"`
}
