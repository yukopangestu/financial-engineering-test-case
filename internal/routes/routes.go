package routes

import (
	"financial-engineering-test-case/internal/config"
	bHandler "financial-engineering-test-case/module/borrower/handler"
	bRepository "financial-engineering-test-case/module/borrower/repository"
	bService "financial-engineering-test-case/module/borrower/service"
	eHandler "financial-engineering-test-case/module/employee/handler"
	eRepository "financial-engineering-test-case/module/employee/repository"
	eService "financial-engineering-test-case/module/employee/service"
	iHandler "financial-engineering-test-case/module/investor/handler"
	iRepository "financial-engineering-test-case/module/investor/repository"
	iService "financial-engineering-test-case/module/investor/service"
	lHandler "financial-engineering-test-case/module/loan/handler"
	lRepository "financial-engineering-test-case/module/loan/repository"
	lService "financial-engineering-test-case/module/loan/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	borrowerRepository := bRepository.NewBorrowerRepository(db)
	borrowerService := bService.NewBorrowerService(borrowerRepository)
	borrowerHandler := bHandler.NewBorrowerHandler(borrowerService)

	employeeRepository := eRepository.NewEmployeeRepository(db)
	employeeService := eService.NewEmployeeService(employeeRepository)
	employeeHandler := eHandler.NewEmployeeHandler(employeeService)

	investorRepository := iRepository.NewInvestorRepository(db)
	investorService := iService.NewInvestorService(investorRepository)
	investorHandler := iHandler.NewInvestorHandler(investorService)

	loanRepository := lRepository.NewLoanRepository(db)
	loanService := lService.NewLoanService(loanRepository, borrowerService, cfg)
	loanHandler := lHandler.NewLoanHandler(loanService)

	v1 := e.Group("/api/v1")
	{
		borrowers := v1.Group("/borrowers")
		borrowers.POST("/", borrowerHandler.CreateBorrower)

		employees := v1.Group("/employees")
		employees.POST("/", employeeHandler.CreateEmployee)

		investors := v1.Group("/investors")
		investors.POST("/", investorHandler.CreateInvestor)

		loans := v1.Group("/loans")
		loans.POST("/propose", loanHandler.ProposeLoan)
		loans.PUT("/{id}/approve", loanHandler.ApproveLoan)
		loans.PUT("/{id}/invest", loanHandler.InvestLoan)
		loans.PUT("/{id}/disburse", loanHandler.DisbursedLoan)
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
