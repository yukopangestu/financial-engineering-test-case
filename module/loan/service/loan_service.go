package service

import (
	bRepository "financial-engineering-test-case/module/borrower/repository"
	"financial-engineering-test-case/module/loan/dto"
	"financial-engineering-test-case/module/loan/repository"
)

type LoanService struct {
	LoanRepository     *repository.LoanRepository
	BorrowerRepository *bRepository.BorrowerRepository
}

func NewLoanService(
	LoanRepository *repository.LoanRepository,
	BorrowerRepository *bRepository.BorrowerRepository,
) *LoanService {
	return &LoanService{
		LoanRepository:     LoanRepository,
		BorrowerRepository: BorrowerRepository,
	}
}

func (s LoanService) ProposeLoan(payload dto.ProposeLoanRequest) {

}
