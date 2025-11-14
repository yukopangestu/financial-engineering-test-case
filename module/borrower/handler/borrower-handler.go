package handler

import (
	"financial-engineering-test-case/module/borrower/domain"
	"financial-engineering-test-case/module/borrower/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BorrowerHandler struct {
	BorrowerService domain.BorrowerService
}

func NewBorrowerHandler(
	BorrowerService domain.BorrowerService) *BorrowerHandler {
	return &BorrowerHandler{
		BorrowerService: BorrowerService,
	}
}

func (h *BorrowerHandler) CreateBorrower(c echo.Context) error {
	var (
		req dto.CreaterBorrower
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	err := h.BorrowerService.CreateNewBorrower(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Borrower successfully created",
	})
}
