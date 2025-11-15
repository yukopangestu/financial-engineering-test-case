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
	result := r.db.Model(&database.Loan{}).Where("id = ?", data.ID).Updates(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r LoanRepository) GetInvestorByID(id uint) (database.Investor, error) {
	var investor database.Investor
	result := r.db.First(&investor, id)
	if result.Error != nil {
		return database.Investor{}, result.Error
	}

	return investor, nil
}

func (r LoanRepository) CreateLoanInvestor(data database.LoanInvestor) error {
	result := r.db.Create(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r LoanRepository) GetLoanInvestorsByLoanID(loanID uint) ([]database.LoanInvestor, error) {
	var loanInvestors []database.LoanInvestor
	result := r.db.Preload("Investor").Where("loan_id = ?", loanID).Find(&loanInvestors)
	if result.Error != nil {
		return nil, result.Error
	}

	return loanInvestors, nil
}
