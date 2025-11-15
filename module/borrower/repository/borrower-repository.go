package repository

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/borrower/domain"
	"financial-engineering-test-case/module/borrower/dto"

	"gorm.io/gorm"
)

type BorrowerRepository struct {
	db *gorm.DB
}

var _ domain.BorrowerRepository = (*BorrowerRepository)(nil)

func NewBorrowerRepository(db *gorm.DB) *BorrowerRepository {
	return &BorrowerRepository{
		db: db,
	}
}

func (b BorrowerRepository) CreateNewBorrower(borrowerData dto.CreaterBorrower) error {
	borrower := database.Borrower{
		Name:           borrowerData.Name,
		BorrowerNumber: borrowerData.BorrowerNum,
		PhoneNumber:    borrowerData.PhoneNumber,
		Email:          borrowerData.Email,
	}

	result := b.db.Create(&borrower)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b BorrowerRepository) GetBorrowers(payload database.Borrower) (database.Borrower, error) {
	borrowers := b.db.Find(&payload)

	if borrowers.Error != nil {
		return database.Borrower{}, borrowers.Error
	}

	return payload, nil
}

func (b BorrowerRepository) GetBorrowerById(id uint) (database.Borrower, error) {
	var borrowerData database.Borrower
	result := b.db.First(&borrowerData, id)
	if result.Error != nil {
		return database.Borrower{}, result.Error
	}

	return borrowerData, nil
}
