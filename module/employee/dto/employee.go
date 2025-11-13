package dto

type CreateEmployee struct {
	Name        string `json:"name" validate:"required"`
	EmployeeNum string `json:"employee_num"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"email"`
}
