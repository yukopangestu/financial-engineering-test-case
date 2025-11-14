package service

import (
	"financial-engineering-test-case/module/borrower/domain"
	"financial-engineering-test-case/module/borrower/dto"
	"fmt"
	"math/rand"
	"time"
)

type BorrowerService struct {
	borrowerRepository domain.BorrowerRepository
}

// Ensure BorrowerService implements domain.BorrowerService interface
var _ domain.BorrowerService = (*BorrowerService)(nil)

func NewBorrowerService(
	borrowerRepository domain.BorrowerRepository) *BorrowerService {
	return &BorrowerService{
		borrowerRepository: borrowerRepository,
	}
}

func (s BorrowerService) CreateNewBorrower(payload dto.CreaterBorrower) error {
	borrowerNumber := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int())
	payload.BorrowerNum = borrowerNumber
	return s.borrowerRepository.CreateNewBorrower(payload)
}
