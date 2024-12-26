package controller

import (
	"DB_Project/src/api/http/request/installment"
	"DB_Project/src/pkg/validation"
	"DB_Project/src/services"
	"DB_Project/src/utils"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v4"
)

var InstallmentNotFound = errors.New("installment not found")
var InstallmentFieldShouldBeUnique = errors.New("installment field should be unique: ")
var InstallmentIdNotSet = errors.New("installment id should be set")

type InstallmentController struct {
	Service *services.InstallmentService
}

func NewInstallmentController() *InstallmentController {
	return &InstallmentController{
		Service: services.NewInstallmentService(),
	}
}

func (controller *InstallmentController) List(c fiber.Ctx) error {
	installments, err := controller.Service.GetInstallments()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"installments": installments,
	})
}

func (controller *InstallmentController) Get(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(InstallmentIdNotSet.Error())
	}

	res, err := controller.Service.GetInstallment(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(InstallmentNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (controller *InstallmentController) Create(c fiber.Ctx) error {
	req := new(installment.CreateInstallmentRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.CreateInstallment(req)
	if err != nil {
		if utils.IsErrorCode(err, "23505") {
			return c.Status(fiber.StatusConflict).SendString(InstallmentFieldShouldBeUnique.Error() + utils.GetErrorConstraintName(err))
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).Send([]byte{})
}

func (controller *InstallmentController) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(InstallmentIdNotSet.Error())
	}

	req := new(installment.UpdateInstallmentRequest)
	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validation.ValidateStruct(req))
	}

	err := controller.Service.UpdateInstallment(req, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).Send([]byte{})
}

func (controller *InstallmentController) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).SendString(InstallmentIdNotSet.Error())
	}

	err := controller.Service.DeleteInstallment(id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).SendString(InstallmentNotFound.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusNoContent).Send([]byte{})
}
