package repository

import (
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/employee/domain"
	"financial-engineering-test-case/module/employee/dto"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

var _ domain.EmployeeRepository = (*EmployeeRepository)(nil)

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (e EmployeeRepository) CreateNewEmployee(employeeData dto.CreateEmployee) error {
	employee := database.Employee{
		Name:           employeeData.Name,
		EmployeeNumber: employeeData.EmployeeNum,
		PhoneNumber:    employeeData.PhoneNumber,
		Email:          employeeData.Email,
	}

	result := e.db.Create(&employee)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
