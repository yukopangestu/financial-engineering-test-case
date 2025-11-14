package service

import (
	"fmt"
	"math/rand"
	"time"

	"financial-engineering-test-case/module/investor/domain"
	"financial-engineering-test-case/module/investor/dto"
)

type InvestorService struct {
	investorRepository domain.InvestorRepository
}

var _ domain.InvestorService = (*InvestorService)(nil)

func NewInvestorService(investorRepository domain.InvestorRepository) *InvestorService {
	return &InvestorService{investorRepository: investorRepository}
}

func (s InvestorService) CreateNewInvestor(payload dto.CreateInvestor) error {
	investorNumber := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int())
	payload.InvestorNum = investorNumber
	return s.investorRepository.CreateNewInvestor(payload)
}
