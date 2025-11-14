package domain

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/investor/dto"
)

type InvestorRepository interface {
	CreateNewInvestor(payload dto.CreateInvestor) error
	GetInvestorById(id uint) (database.Investor, error)
}

type InvestorService interface {
	CreateNewInvestor(payload dto.CreateInvestor) error
}
