package controller

import (
	"DB_Project/src/api/http/exception"
	"DB_Project/src/api/http/request/installment"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"github.com/gofiber/fiber/v3"
)

type InstallmentController struct {
	service          *services.InstallmentService
	exceptionHandler exception.Exception
}

func NewInstallmentController() *InstallmentController {
	return &InstallmentController{
		service:          services.NewInstallmentService(),
		exceptionHandler: exception.NewInstallmentExceptions(),
	}
}

func (controller *InstallmentController) List(c fiber.Ctx) error {
	installments, err := controller.service.GetInstallments()
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"installments": installments,
	})
}

func (controller *InstallmentController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.InstallmentIdNotSet.Error())
	}

	res, err := controller.service.GetInstallment(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *InstallmentController) Create(c fiber.Ctx) error {
	req := new(installment.CreateInstallmentRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.CreateInstallment(req)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *InstallmentController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.InstallmentIdNotSet.Error())
	}

	req := new(installment.UpdateInstallmentRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.service.UpdateInstallment(req, id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}

	return c.Status(fiber.StatusOK).Send([]byte{})
}

func (controller *InstallmentController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(exception.InstallmentIdNotSet.Error())
	}

	err := controller.service.DeleteInstallment(id)
	if err != nil {
		return controller.exceptionHandler.Handle(err, c)
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
