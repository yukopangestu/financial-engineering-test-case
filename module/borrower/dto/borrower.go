package dto

type CreaterBorrower struct {
	Name        string `json:"name" validate:"required"`
	BorrowerNum string `json:"borrower_num"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"email"`
}
