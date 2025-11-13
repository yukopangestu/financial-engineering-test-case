package handler

import (
	"financial-engineering-test-case/module/employee/dto"
	"financial-engineering-test-case/module/employee/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeHandler(
	EmployeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: EmployeeService,
	}
}

func (h *EmployeeHandler) CreateEmployee(c echo.Context) error {
	var (
		req dto.CreateEmployee
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	err := h.EmployeeService.CreateNewEmployee(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Employee successfully created",
	})
}
