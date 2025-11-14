package domain

import "financial-engineering-test-case/module/employee/dto"

type EmployeeRepository interface {
	CreateNewEmployee(employeeData dto.CreateEmployee) error
}

type EmployeeService interface {
	CreateNewEmployee(employeeData dto.CreateEmployee) error
}
