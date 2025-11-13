package service

import (
	"financial-engineering-test-case/module/employee/dto"
	"financial-engineering-test-case/module/employee/repository"
	"fmt"
	"math/rand"
	"time"
)

type EmployeeService struct {
	employeeRepository *repository.EmployeeRepository
}

func NewEmployeeService(
	employeeRepository *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (s EmployeeService) CreateNewEmployee(payload dto.CreateEmployee) error {
	employeeNumber := fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int())
	payload.EmployeeNum = employeeNumber
	return s.employeeRepository.CreateNewEmployee(payload)
}
