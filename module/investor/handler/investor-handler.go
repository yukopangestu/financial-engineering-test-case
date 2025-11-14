package handler

import (
	"net/http"

	"financial-engineering-test-case/module/investor/domain"
	"financial-engineering-test-case/module/investor/dto"

	"github.com/labstack/echo/v4"
)

type InvestorHandler struct {
	InvestorService domain.InvestorService
}

func NewInvestorHandler(investorService domain.InvestorService) *InvestorHandler {
	return &InvestorHandler{InvestorService: investorService}
}

func (h *InvestorHandler) CreateInvestor(c echo.Context) error {
	var req dto.CreateInvestor

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	err := h.InvestorService.CreateNewInvestor(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Investor successfully created",
	})
}
