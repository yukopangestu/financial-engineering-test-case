package handler

import (
	"financial-engineering-test-case/module/loan/dto"
	"financial-engineering-test-case/module/loan/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	LoanService *service.LoanService
}

func NewLoanHandler(
	LoanService *service.LoanService,
) *LoanHandler {
	return &LoanHandler{
		LoanService: LoanService,
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

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.LoanService.ProposeLoan(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Loan proposal received",
	})
}

func (h *LoanHandler) ApproveLoan(c echo.Context) error {
	var req dto.ApproveLoanRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	id := c.Param("id")
	LoanID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error when Converting Uid",
		})
	}

	req.ID = uint(LoanID)
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Loan successfully approved",
	})
}

func (h *LoanHandler) InvestLoan(c echo.Context) error {
	var req dto.InvestLoanRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id := c.Param("id")
	loanID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid loan ID",
		})
	}

	err = h.LoanService.InvestLoan(&req, uint(loanID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Loan successfully invested and agreement letter generated",
	})
}

func (h *LoanHandler) DisbursedLoan(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Loan successfully invested and agreement letter generated",
	})
}
