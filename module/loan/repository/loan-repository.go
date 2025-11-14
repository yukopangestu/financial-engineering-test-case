package repository

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/loan/domain"

	"gorm.io/gorm"
)

type LoanRepository struct {
	db *gorm.DB
}

var _ domain.LoanRepository = (*LoanRepository)(nil)

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

func (r LoanRepository) GetLoanByID(id uint) (database.Loan, error) {
	var loan database.Loan
	result := r.db.First(&loan, id)
	if result.Error != nil {
		return database.Loan{}, result.Error
	}

	return loan, nil
}

func (r LoanRepository) UploadLoanByID(data database.Loan) error {
	result := r.db.Save(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
