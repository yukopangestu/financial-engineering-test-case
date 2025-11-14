package routes

import (
	bHandler "financial-engineering-test-case/module/borrower/handler"
	bRepository "financial-engineering-test-case/module/borrower/repository"
	bService "financial-engineering-test-case/module/borrower/service"
	eHandler "financial-engineering-test-case/module/employee/handler"
	eRepository "financial-engineering-test-case/module/employee/repository"
	eService "financial-engineering-test-case/module/employee/service"
	lHandler "financial-engineering-test-case/module/loan/handler"
	lRepository "financial-engineering-test-case/module/loan/repository"
	lService "financial-engineering-test-case/module/loan/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize handlers
	borrowerRepository := bRepository.NewBorrowerRepository(db)
	borrowerService := bService.NewBorrowerService(borrowerRepository)
	borrowerHandler := bHandler.NewBorrowerHandler(borrowerService)

	employeeRepository := eRepository.NewEmployeeRepository(db)
	employeeService := eService.NewEmployeeService(employeeRepository)
	employeeHandler := eHandler.NewEmployeeHandler(employeeService)

	loanRepository := lRepository.NewLoanRepository(db)
	loanService := lService.NewLoanService(loanRepository, borrowerService)
	loanHandler := lHandler.NewLoanHandler(loanService)

	// API v1 routes
	v1 := e.Group("/api/v1")
	{
		// Managing Borrowers data
		borrowers := v1.Group("/borrowers")
		borrowers.POST("/", borrowerHandler.CreateBorrower)

		// Managing Employees data
		employees := v1.Group("/employees")
		employees.POST("/", employeeHandler.CreateEmployee)

		// Loan routes
		loans := v1.Group("/loans")
		loans.POST("/propose", loanHandler.ProposeLoan)
		loans.PUT("/{id}/approve", loanHandler.ApproveLoan)
		loans.PUT("/{id}/invest", loanHandler.InvestLoan)
		loans.PUT("/{id}/disburse", loanHandler.InvestLoan)
	}

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
