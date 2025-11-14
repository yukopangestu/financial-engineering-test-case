package repository

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/investor/domain"
	"financial-engineering-test-case/module/investor/dto"
	"gorm.io/gorm"
)

type InvestorRepository struct {
	db *gorm.DB
}

var _ domain.InvestorRepository = (*InvestorRepository)(nil)

func NewInvestorRepository(db *gorm.DB) *InvestorRepository {
	return &InvestorRepository{db: db}
}

func (r InvestorRepository) CreateNewInvestor(data dto.CreateInvestor) error {
	investor := database.Investor{
		Name:           data.Name,
		InvestorNumber: data.InvestorNum,
		PhoneNumber:    data.PhoneNumber,
		Email:          data.Email,
		InvestedAmount: 0,
	}

	result := r.db.Create(&investor)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r InvestorRepository) GetInvestorById(id uint) (database.Investor, error) {
	var investor database.Investor
	result := r.db.First(&investor, id)
	if result.Error != nil {
		return database.Investor{}, result.Error
	}
	return investor, nil
}
