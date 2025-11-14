package domain

import "financial-engineering-test-case/module/borrower/dto"

type BorrowerRepository interface {
	CreateNewBorrower(payload dto.CreaterBorrower) error
}

type BorrowerService interface {
	CreateNewBorrower(payload dto.CreaterBorrower) error
}
