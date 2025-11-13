package service

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/borrower/dto"
	"financial-engineering-test-case/module/borrower/repository"
	"fmt"
	"math/rand"
	"strconv"
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

	borrowerNumber := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int(20000))
	payload.BorrowerNum = borrowerNumber
	return s.borrowerRepository.CreateNewBorrower(payload)
}

func (s BorrowerService) GetBorrowers(payload database.Borrower) (database.Borrower, error) {
	data, err := s.borrowerRepository.GetBorrowers(payload)

	if err != nil {
		return database.Borrower{}, err
	}

	return data, nil
}
