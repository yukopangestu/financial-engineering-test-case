package service

import (
	"errors"
	"financial-engineering-test-case/internal/database"
	bService "financial-engineering-test-case/module/borrower/service"
	"financial-engineering-test-case/module/loan/dto"
	"financial-engineering-test-case/module/loan/repository"
	"fmt"
	"math/rand"
	"time"
)

type LoanService struct {
	LoanRepository  *repository.LoanRepository
	BorrowerService *bService.BorrowerService
}

func NewLoanService(
	LoanRepository *repository.LoanRepository,
	BorrowerService *bService.BorrowerService,
) *LoanService {
	return &LoanService{
		LoanRepository:  LoanRepository,
		BorrowerService: BorrowerService,
	}
}

func (s LoanService) ProposeLoan(payload dto.ProposeLoanRequest) error {
	var data database.Loan

	borrowers, err := s.BorrowerService.GetBorrowerById(payload.BorrowerId)
	if err != nil {
		return fmt.Errorf("Error while fecthing the borrowers", err)
	}
	if borrowers.ID == 0 {
		return errors.New("borrower not exist")
	}

	data = database.Loan{
		Amount:     payload.Amount,
		BorrowerId: payload.BorrowerId,
		Interest:   payload.InterestRate,
		LoanNumber: fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int()),
	}

	err = s.LoanRepository.CreateNewLoan(data)
	if err != nil {
		return fmt.Errorf("Error when creating new loan", err)
	}

	return nil
}
