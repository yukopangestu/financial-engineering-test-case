package dto

type CreateInvestor struct {
	Name        string `json:"name" validate:"required"`
	InvestorNum string `json:"investor_num"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"email"`
}
