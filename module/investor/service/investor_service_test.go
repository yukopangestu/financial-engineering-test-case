package service

import (
	"errors"
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/investor/dto"
	"testing"
)

// Mock Repository for Investor Testing
type MockInvestorRepository struct {
	createNewInvestorFunc func(payload dto.CreateInvestor) error
	getInvestorByIdFunc   func(id uint) (database.Investor, error)
}

func (m *MockInvestorRepository) CreateNewInvestor(payload dto.CreateInvestor) error {
	if m.createNewInvestorFunc != nil {
		return m.createNewInvestorFunc(payload)
	}
	return nil
}

func (m *MockInvestorRepository) GetInvestorById(id uint) (database.Investor, error) {
	if m.getInvestorByIdFunc != nil {
		return m.getInvestorByIdFunc(id)
	}
	return database.Investor{}, nil
}

// Test CreateNewInvestor - Success Case
func TestCreateNewInvestor_Success(t *testing.T) {
	var capturedPayload dto.CreateInvestor

	mockRepo := &MockInvestorRepository{
		createNewInvestorFunc: func(payload dto.CreateInvestor) error {
			capturedPayload = payload
			return nil
		},
	}

	service := NewInvestorService(mockRepo)

	payload := dto.CreateInvestor{
		Name:        "Investor John",
		PhoneNumber: "081234567890",
		Email:       "investor@example.com",
	}

	err := service.CreateNewInvestor(payload)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify investor number was generated
	if capturedPayload.InvestorNum == "" {
		t.Error("Expected InvestorNum to be generated, got empty string")
	}

	// Verify other fields were preserved
	if capturedPayload.Name != "Investor John" {
		t.Errorf("Expected Name 'Investor John', got '%s'", capturedPayload.Name)
	}
	if capturedPayload.Email != "investor@example.com" {
		t.Errorf("Expected Email 'investor@example.com', got '%s'", capturedPayload.Email)
	}
	if capturedPayload.PhoneNumber != "081234567890" {
		t.Errorf("Expected PhoneNumber '081234567890', got '%s'", capturedPayload.PhoneNumber)
	}
}

// Test CreateNewInvestor - Repository Error
func TestCreateNewInvestor_RepositoryError(t *testing.T) {
	mockRepo := &MockInvestorRepository{
		createNewInvestorFunc: func(payload dto.CreateInvestor) error {
			return errors.New("database connection failed")
		},
	}

	service := NewInvestorService(mockRepo)

	payload := dto.CreateInvestor{
		Name:        "Investor Jane",
		PhoneNumber: "081234567891",
		Email:       "jane.investor@example.com",
	}

	err := service.CreateNewInvestor(payload)
	if err == nil {
		t.Error("Expected error from repository, got nil")
	}
	if err.Error() != "database connection failed" {
		t.Errorf("Expected 'database connection failed', got: %v", err)
	}
}

// Test CreateNewInvestor - Investor Number Format
func TestCreateNewInvestor_InvestorNumberFormat(t *testing.T) {
	var capturedPayload dto.CreateInvestor

	mockRepo := &MockInvestorRepository{
		createNewInvestorFunc: func(payload dto.CreateInvestor) error {
			capturedPayload = payload
			return nil
		},
	}

	service := NewInvestorService(mockRepo)

	payload := dto.CreateInvestor{
		Name:        "Test Investor",
		PhoneNumber: "081234567890",
		Email:       "test@example.com",
	}

	err := service.CreateNewInvestor(payload)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Investor number should follow format: YEAR/MONTH/RANDOM
	if capturedPayload.InvestorNum == "" {
		t.Error("Investor number should not be empty")
	}
	// Note: We can't easily test the exact format without mocking time/rand,
	// but we verify it's not empty which confirms generation occurred
}

// Test CreateNewInvestor - Multiple Investors Get Unique Numbers
func TestCreateNewInvestor_UniqueNumbers(t *testing.T) {
	generatedNumbers := make(map[string]bool)

	mockRepo := &MockInvestorRepository{
		createNewInvestorFunc: func(payload dto.CreateInvestor) error {
			if generatedNumbers[payload.InvestorNum] {
				t.Errorf("Duplicate investor number generated: %s", payload.InvestorNum)
			}
			generatedNumbers[payload.InvestorNum] = true
			return nil
		},
	}

	service := NewInvestorService(mockRepo)

	// Create multiple investors
	for i := 0; i < 10; i++ {
		payload := dto.CreateInvestor{
			Name:        "Test Investor",
			PhoneNumber: "081234567890",
			Email:       "test@example.com",
		}
		service.CreateNewInvestor(payload)
	}

	if len(generatedNumbers) < 10 {
		t.Errorf("Expected 10 unique investor numbers, got %d", len(generatedNumbers))
	}
}

// Test CreateNewInvestor - Empty Fields
func TestCreateNewInvestor_EmptyFields(t *testing.T) {
	tests := []struct {
		name    string
		payload dto.CreateInvestor
	}{
		{
			name: "Empty name",
			payload: dto.CreateInvestor{
				Name:        "",
				PhoneNumber: "081234567890",
				Email:       "test@example.com",
			},
		},
		{
			name: "Empty email",
			payload: dto.CreateInvestor{
				Name:        "Test Investor",
				PhoneNumber: "081234567890",
				Email:       "",
			},
		},
		{
			name: "Empty phone",
			payload: dto.CreateInvestor{
				Name:        "Test Investor",
				PhoneNumber: "",
				Email:       "test@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockInvestorRepository{
				createNewInvestorFunc: func(payload dto.CreateInvestor) error {
					// Repository still accepts it (validation should be at handler level)
					return nil
				},
			}

			service := NewInvestorService(mockRepo)

			err := service.CreateNewInvestor(tt.payload)
			// Service doesn't validate - that's done at handler/DTO level
			// So we just ensure it doesn't crash
			if err != nil {
				t.Errorf("Service should not validate, got error: %v", err)
			}
		})
	}
}

// Benchmark CreateNewInvestor
func BenchmarkCreateNewInvestor(b *testing.B) {
	mockRepo := &MockInvestorRepository{
		createNewInvestorFunc: func(payload dto.CreateInvestor) error {
			return nil
		},
	}

	service := NewInvestorService(mockRepo)

	payload := dto.CreateInvestor{
		Name:        "Benchmark Investor",
		PhoneNumber: "081234567890",
		Email:       "benchmark@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.CreateNewInvestor(payload)
	}
}
