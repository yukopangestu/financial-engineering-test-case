package domain

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/loan/dto"
)

type LoanRepository interface {
	CreateNewLoan(data database.Loan) error
	GetLoanByID(id uint) (database.Loan, error)
	UploadLoanByID(data database.Loan) error
}

type LoanService interface {
	ProposeLoan(payload *dto.ProposeLoanRequest) error
	ApproveLoan(payload *dto.ApproveLoanRequest) error
	InvestLoan(payload *dto.InvestLoanRequest, id uint) error
	DisburseLoan(payload *dto.DisbursedLoanRequest, id uint) error
}
