package services

import (
	"DB_Project/src/api/http/request/loan"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"time"
)

type LoanService struct {
	Repository *repositories.LoanRepository
}

func NewLoanService() *LoanService {
	return &LoanService{
		Repository: repositories.NewLoanRepository(),
	}
}

func (service *LoanService) GetLoans() ([]*models.Loan, error) {
	return service.Repository.List()
}

func (service *LoanService) GetLoan(id string) (*models.Loan, error) {
	return service.Repository.Get(id)
}

func (service *LoanService) CreateLoan(req *loan.CreateLoanRequest) error {
	finishAt, _ := time.Parse("2006-01-02", req.FinishedAt)

	// Map the request to the Loan model
	loan := &models.Loan{
		CustomerID:      req.CustomerID,
		Type:            req.Type,
		Status:          req.Status,
		Amount:          req.Amount,
		InterestRate:    req.InterestRate,
		RepaymentPeriod: req.RepaymentPeriod,
		FinishedAt:      finishAt,
	}

	return service.Repository.Create(loan)
}

func (service *LoanService) UpdateLoan(req *loan.UpdateLoanRequest, id string) error {
	if req.CustomerID != nil {
		err := service.Repository.UpdateField("customer_id", id, *req.CustomerID)
		if err != nil {
			return err
		}
	}

	if req.Type != nil && *req.Type != "" {
		err := service.Repository.UpdateField("type", id, *req.Type)
		if err != nil {
			return err
		}
	}

	if req.Status != nil && *req.Status != "" {
		err := service.Repository.UpdateField("status", id, *req.Status)
		if err != nil {
			return err
		}
	}

	if req.Amount != nil {
		err := service.Repository.UpdateField("amount", id, *req.Amount)
		if err != nil {
			return err
		}
	}

	if req.InterestRate != nil {
		err := service.Repository.UpdateField("interest_rate", id, *req.InterestRate)
		if err != nil {
			return err
		}
	}

	if req.RepaymentPeriod != nil {
		err := service.Repository.UpdateField("repayment_period", id, *req.RepaymentPeriod)
		if err != nil {
			return err
		}
	}

	if req.FinishedAt != nil {
		err := service.Repository.UpdateField("finished_at", id, *req.FinishedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *LoanService) DeleteLoan(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
