package repository

import (
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
	result := e.db.Create(&employeeData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
