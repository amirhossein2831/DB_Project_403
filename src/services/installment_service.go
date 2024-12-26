package services

import (
	"DB_Project/src/api/http/request/installment"
	"DB_Project/src/models"
	"DB_Project/src/repositories"
	"time"
)

type InstallmentService struct {
	Repository *repositories.InstallmentRepository
}

func NewInstallmentService() *InstallmentService {
	return &InstallmentService{
		Repository: repositories.NewInstallmentRepository(),
	}
}

func (service *InstallmentService) GetInstallments() ([]*models.Installment, error) {
	return service.Repository.List()
}

func (service *InstallmentService) GetInstallment(id string) (*models.Installment, error) {
	return service.Repository.Get(id)
}

func (service *InstallmentService) CreateInstallment(req *installment.CreateInstallmentRequest) error {
	var paidDate time.Time
	dueDate, _ := time.Parse("2006-01-02", req.DueDate)
	if req.PaidDate != nil {
		paidDate, _ = time.Parse("2006-01-02", *req.PaidDate)

	}

	installment := &models.Installment{
		LoanID:       req.LoanID,
		AmountPaid:   req.AmountPaid,
		InterestPaid: req.InterestPaid,
		TotalPaid:    req.TotalPaid,
		DueDate:      dueDate,
		PaidDate:     &paidDate,
	}

	return service.Repository.Create(installment)
}

func (service *InstallmentService) UpdateInstallment(req *installment.UpdateInstallmentRequest, id string) error {
	if req.AmountPaid != nil {
		err := service.Repository.UpdateField("amount_paid", id, *req.AmountPaid)
		if err != nil {
			return err
		}
	}

	if req.InterestPaid != nil {
		err := service.Repository.UpdateField("interest_paid", id, *req.InterestPaid)
		if err != nil {
			return err
		}
	}

	if req.TotalPaid != nil {
		err := service.Repository.UpdateField("total_paid", id, *req.TotalPaid)
		if err != nil {
			return err
		}
	}

	if req.DueDate != nil {
		err := service.Repository.UpdateField("due_date", id, *req.DueDate)
		if err != nil {
			return err
		}
	}

	err := service.Repository.UpdateField("paid_date", id, req.PaidDate)
	if err != nil {
		return err
	}

	return nil
}

func (service *InstallmentService) DeleteInstallment(id string) error {
	_, err := service.Repository.Get(id)
	if err != nil {
		return err
	}

	return service.Repository.Delete(id)
}
