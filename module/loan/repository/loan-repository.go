package repository

import (
	"financial-engineering-test-case/internal/database"

	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(
	db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		db: db,
	}
}

func (r LoanRepository) CreateNewLoan(data database.Loan) error {
	result := r.db.Create(&data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
