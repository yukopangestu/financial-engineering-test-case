package service

import (
	"financial-engineering-test-case/module/borrower/dto"
	"financial-engineering-test-case/module/borrower/repository"
	"fmt"
	"math/rand"
	"time"
)

type BorrowerService struct {
	borrowerRepository *repository.BorrowerRepository
}

func NewBorrowerService(
	borrowerRepository *repository.BorrowerRepository) *BorrowerService {
	return &BorrowerService{
		borrowerRepository: borrowerRepository,
	}
}

func (s BorrowerService) CreateNewBorrower(payload dto.CreaterBorrower) error {
	borrowerNumber := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int())
	payload.BorrowerNum = borrowerNumber
	return s.borrowerRepository.CreateNewBorrower(payload)
}
