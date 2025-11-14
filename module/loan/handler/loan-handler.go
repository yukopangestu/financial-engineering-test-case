package handler

import (
	"financial-engineering-test-case/module/loan/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	// Add your service dependency here
	// loanService service.LoanService
}

func NewLoanHandler() *LoanHandler {
	return &LoanHandler{
		// Initialize your service here
	}
}

func (h *LoanHandler) ProposeLoan(c echo.Context) error {
	var req dto.ProposeLoanRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	// Validate request
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call service layer
	// result, err := h.loanService.ProposeLoan(c.Request().Context(), &req)
	// if err != nil {
	//     return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Loan proposal received",
	})
}
