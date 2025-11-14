package dto

type ProposeLoanRequest struct {
	BorrowerId   string  `json:"borrower_id" validate:"required"`
	Amount       float64 `json:"amount" validate:"required, gt=0"`
	InterestRate float64 `json:"interest_rate" validate:"required, gt=0"`
}
