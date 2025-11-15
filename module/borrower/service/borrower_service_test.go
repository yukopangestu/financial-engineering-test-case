package service

import (
	"errors"
	"financial-engineering-test-case/internal/database"
	"financial-engineering-test-case/module/borrower/dto"
	"testing"
)

// Mock Repository for Borrower Testing
type MockBorrowerRepository struct {
	createNewBorrowerFunc func(payload dto.CreaterBorrower) error
	getBorrowerByIdFunc   func(id uint) (database.Borrower, error)
}

func (m *MockBorrowerRepository) CreateNewBorrower(payload dto.CreaterBorrower) error {
	if m.createNewBorrowerFunc != nil {
		return m.createNewBorrowerFunc(payload)
	}
	return nil
}

func (m *MockBorrowerRepository) GetBorrowerById(id uint) (database.Borrower, error) {
	if m.getBorrowerByIdFunc != nil {
		return m.getBorrowerByIdFunc(id)
	}
	return database.Borrower{}, nil
}

// Test CreateNewBorrower - Success Case
func TestCreateNewBorrower_Success(t *testing.T) {
	var capturedPayload dto.CreaterBorrower

	mockRepo := &MockBorrowerRepository{
		createNewBorrowerFunc: func(payload dto.CreaterBorrower) error {
			capturedPayload = payload
			return nil
		},
	}

	service := NewBorrowerService(mockRepo)

	payload := dto.CreaterBorrower{
		Name:        "John Doe",
		PhoneNumber: "081234567890",
		Email:       "john@example.com",
	}

	err := service.CreateNewBorrower(payload)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify borrower number was generated
	if capturedPayload.BorrowerNum == "" {
		t.Error("Expected BorrowerNum to be generated, got empty string")
	}

	// Verify other fields were preserved
	if capturedPayload.Name != "John Doe" {
		t.Errorf("Expected Name 'John Doe', got '%s'", capturedPayload.Name)
	}
	if capturedPayload.Email != "john@example.com" {
		t.Errorf("Expected Email 'john@example.com', got '%s'", capturedPayload.Email)
	}
	if capturedPayload.PhoneNumber != "081234567890" {
		t.Errorf("Expected PhoneNumber '081234567890', got '%s'", capturedPayload.PhoneNumber)
	}
}

// Test CreateNewBorrower - Repository Error
func TestCreateNewBorrower_RepositoryError(t *testing.T) {
	mockRepo := &MockBorrowerRepository{
		createNewBorrowerFunc: func(payload dto.CreaterBorrower) error {
			return errors.New("database error")
		},
	}

	service := NewBorrowerService(mockRepo)

	payload := dto.CreaterBorrower{
		Name:        "Jane Doe",
		PhoneNumber: "081234567891",
		Email:       "jane@example.com",
	}

	err := service.CreateNewBorrower(payload)
	if err == nil {
		t.Error("Expected error from repository, got nil")
	}
	if err.Error() != "database error" {
		t.Errorf("Expected 'database error', got: %v", err)
	}
}

// Test CreateNewBorrower - Borrower Number Generation Uniqueness
func TestCreateNewBorrower_BorrowerNumberGeneration(t *testing.T) {
	generatedNumbers := make(map[string]bool)

	mockRepo := &MockBorrowerRepository{
		createNewBorrowerFunc: func(payload dto.CreaterBorrower) error {
			if generatedNumbers[payload.BorrowerNum] {
				t.Errorf("Duplicate borrower number generated: %s", payload.BorrowerNum)
			}
			generatedNumbers[payload.BorrowerNum] = true
			return nil
		},
	}

	service := NewBorrowerService(mockRepo)

	// Create multiple borrowers and ensure unique numbers (though rand.Int makes this probabilistic)
	for i := 0; i < 10; i++ {
		payload := dto.CreaterBorrower{
			Name:        "Test User",
			PhoneNumber: "081234567890",
			Email:       "test@example.com",
		}
		service.CreateNewBorrower(payload)
	}

	if len(generatedNumbers) < 10 {
		t.Errorf("Expected 10 unique borrower numbers, got %d", len(generatedNumbers))
	}
}

// Test GetBorrowerById - Success
func TestGetBorrowerById_Success(t *testing.T) {
	expectedBorrower := database.Borrower{
		ID:             1,
		BorrowerNumber: "2024/1/12345",
		Name:           "John Doe",
		PhoneNumber:    "081234567890",
		Email:          "john@example.com",
	}

	mockRepo := &MockBorrowerRepository{
		getBorrowerByIdFunc: func(id uint) (database.Borrower, error) {
			if id == 1 {
				return expectedBorrower, nil
			}
			return database.Borrower{}, errors.New("borrower not found")
		},
	}

	service := NewBorrowerService(mockRepo)

	borrower, err := service.GetBorrowerById(1)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if borrower.ID != expectedBorrower.ID {
		t.Errorf("Expected ID %d, got %d", expectedBorrower.ID, borrower.ID)
	}
	if borrower.Name != expectedBorrower.Name {
		t.Errorf("Expected Name '%s', got '%s'", expectedBorrower.Name, borrower.Name)
	}
	if borrower.Email != expectedBorrower.Email {
		t.Errorf("Expected Email '%s', got '%s'", expectedBorrower.Email, borrower.Email)
	}
}

// Test GetBorrowerById - Not Found
func TestGetBorrowerById_NotFound(t *testing.T) {
	mockRepo := &MockBorrowerRepository{
		getBorrowerByIdFunc: func(id uint) (database.Borrower, error) {
			return database.Borrower{}, errors.New("borrower not found")
		},
	}

	service := NewBorrowerService(mockRepo)

	_, err := service.GetBorrowerById(999)
	if err == nil {
		t.Error("Expected error for non-existent borrower, got nil")
	}
}

// Test GetBorrowerById - Repository Error
func TestGetBorrowerById_RepositoryError(t *testing.T) {
	mockRepo := &MockBorrowerRepository{
		getBorrowerByIdFunc: func(id uint) (database.Borrower, error) {
			return database.Borrower{}, errors.New("database connection error")
		},
	}

	service := NewBorrowerService(mockRepo)

	_, err := service.GetBorrowerById(1)
	if err == nil {
		t.Error("Expected database error, got nil")
	}
}

// Benchmark CreateNewBorrower
func BenchmarkCreateNewBorrower(b *testing.B) {
	mockRepo := &MockBorrowerRepository{
		createNewBorrowerFunc: func(payload dto.CreaterBorrower) error {
			return nil
		},
	}

	service := NewBorrowerService(mockRepo)

	payload := dto.CreaterBorrower{
		Name:        "Benchmark User",
		PhoneNumber: "081234567890",
		Email:       "benchmark@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.CreateNewBorrower(payload)
	}
}

// Benchmark GetBorrowerById
func BenchmarkGetBorrowerById(b *testing.B) {
	mockRepo := &MockBorrowerRepository{
		getBorrowerByIdFunc: func(id uint) (database.Borrower, error) {
			return database.Borrower{
				ID:    id,
				Name:  "Test User",
				Email: "test@example.com",
			}, nil
		},
	}

	service := NewBorrowerService(mockRepo)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.GetBorrowerById(1)
	}
}
