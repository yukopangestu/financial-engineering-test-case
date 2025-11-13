package dto

type ProposeLoanRequest struct {
	BorrowerId   string  `json:"borrower_id"`
	Amount       float64 `json:"amount"`
	InterestRate float64 `json:"interest_rate"`
}
