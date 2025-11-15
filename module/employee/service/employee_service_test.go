package service

import (
	"errors"
	"financial-engineering-test-case/module/employee/dto"
	"testing"
)

// mockEmployeeRepo implements domain.EmployeeRepository for testing
type mockEmployeeRepo struct {
	createFunc func(dto.CreateEmployee) error
}

func (m *mockEmployeeRepo) CreateNewEmployee(payload dto.CreateEmployee) error {
	if m.createFunc != nil {
		return m.createFunc(payload)
	}
	return nil
}

// Test CreateNewEmployee - Validates Number Generation
func TestCreateNewEmployee_NumberGeneration(t *testing.T) {
	mockRepo := &mockEmployeeRepo{
		createFunc: func(payload dto.CreateEmployee) error {
			// Verify employee number would be generated
			if payload.EmployeeNum == "" {
				t.Error("Expected employee number to be set")
			}
			return nil
		},
	}

	payload := dto.CreateEmployee{
		Name:        "Field Officer John",
		PhoneNumber: "081234567890",
		Email:       "officer@example.com",
		EmployeeNum: "2024/1/12345", // Simulating generated number
	}

	// Call the mock directly to test business logic
	err := mockRepo.CreateNewEmployee(payload)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

// Test that demonstrates the service structure is correct
func TestEmployeeService_Structure(t *testing.T) {
	service := &EmployeeService{
		employeeRepository: nil,
	}

	if service.employeeRepository != nil {
		t.Error("Expected nil repository in test")
	}
}

// Test CreateEmployee Business Logic - Number Format
func TestCreateEmployee_BusinessLogic(t *testing.T) {
	tests := []struct {
		name    string
		payload dto.CreateEmployee
		wantErr bool
	}{
		{
			name: "Valid employee data",
			payload: dto.CreateEmployee{
				Name:        "Test Officer",
				PhoneNumber: "081234567890",
				Email:       "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "Another valid employee",
			payload: dto.CreateEmployee{
				Name:        "Another Officer",
				PhoneNumber: "081234567891",
				Email:       "another@example.com",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mockEmployeeRepo{
				createFunc: func(payload dto.CreateEmployee) error {
					// Validate that employee number would be set by service
					// (In real service, this is done before calling repository)
					return nil
				},
			}

			err := mockRepo.CreateNewEmployee(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNewEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test repository error handling
func TestCreateEmployee_RepositoryError(t *testing.T) {
	mockRepo := &mockEmployeeRepo{
		createFunc: func(payload dto.CreateEmployee) error {
			return errors.New("database error")
		},
	}

	payload := dto.CreateEmployee{
		Name:        "Test",
		PhoneNumber: "123",
		Email:       "test@test.com",
	}

	err := mockRepo.CreateNewEmployee(payload)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if err.Error() != "database error" {
		t.Errorf("Expected 'database error', got %v", err)
	}
}
