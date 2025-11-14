package domain

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/borrower/dto"
)

type BorrowerRepository interface {
	CreateNewBorrower(payload dto.CreaterBorrower) error
	GetBorrowerById(id uint) (database.Borrower, error)
}

type BorrowerService interface {
	CreateNewBorrower(payload dto.CreaterBorrower) error
}
