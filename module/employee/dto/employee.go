package dto

type Borrower struct {
	Id          string `json:"id"`
	EmployeeNo  string `json:"employee_no"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
