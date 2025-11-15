package dto

import "mime/multipart"

type ProposeLoanRequest struct {
	BorrowerId   uint    `json:"borrower_id" validate:"required"`
	Amount       float64 `json:"amount" validate:"required, gt=0"`
	InterestRate float64 `json:"interest_rate" validate:"required, gt=0"`
}

type ApproveLoanRequest struct {
	ID           uint                  `json:"id"`
	EmployeeID   uint                  `json:"employee_id" validate:"required"`
	PhotoOfVisit *multipart.FileHeader `from:"document" validate:"required"`
	ApprovalDate string                `json:"approval_date" validate:"required"`
}

type InvestLoanRequest struct {
	Investors []Investor `json:"investors" validate:"required"`
}

type Investor struct {
	ID               uint    `json:"id" validate:"required"`
	InvestmentAmount float64 `json:"investment_amount" validate:"required,gt=0"`
}

type DisbursedLoanRequest struct {
	EmployeeID            uint                  `json:"employee_id" validate:"required"`
	SignedAgreementLetter *multipart.FileHeader `form:"signed_agreement_letter" validate:"required"`
	DisbursementDate      string                `json:"disbursement_date" validate:"required"`
}
