package routes

import (
	"financial-engineering-test-case/module/loan/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize handlers
	loanHandler := handler.NewLoanHandler()

	// API v1 routes
	v1 := e.Group("/api/v1")
	{
		// Loan routes
		loans := v1.Group("/loans")
		loans.POST("/propose", loanHandler.ProposeLoan)
		// Add more loan endpoints here
		// loans.GET("/:id", loanHandler.GetLoan)
		// loans.GET("", loanHandler.ListLoans)
	}

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
