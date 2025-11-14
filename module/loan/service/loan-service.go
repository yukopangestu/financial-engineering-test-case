package service

import (
	"errors"
	"financial-engineering-test-case/internal/database"
	bService "financial-engineering-test-case/module/borrower/service"
	"financial-engineering-test-case/module/loan/domain"
	"financial-engineering-test-case/module/loan/dto"
	"financial-engineering-test-case/module/loan/repository"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type LoanService struct {
	LoanRepository  *repository.LoanRepository
	BorrowerService *bService.BorrowerService
}

var _ domain.LoanService = (*LoanService)(nil)

func NewLoanService(
	LoanRepository *repository.LoanRepository,
	BorrowerService *bService.BorrowerService,
) *LoanService {
	return &LoanService{
		LoanRepository:  LoanRepository,
		BorrowerService: BorrowerService,
	}
}

func (s LoanService) ProposeLoan(payload *dto.ProposeLoanRequest) error {
	var data database.Loan

	borrowers, err := s.BorrowerService.GetBorrowerById(payload.BorrowerId)
	if err != nil {
		return fmt.Errorf("Error while fecthing the borrowers", err)
	}
	if borrowers.ID == 0 {
		return errors.New("borrower not exist")
	}

	data = database.Loan{
		Amount:     payload.Amount,
		BorrowerID: payload.BorrowerId,
		Interest:   payload.InterestRate,
		LoanNumber: fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Month(), rand.Int()),
	}

	err = s.LoanRepository.CreateNewLoan(data)
	if err != nil {
		return fmt.Errorf("Error when creating new loan", err)
	}

	return nil
}

func (s LoanService) ApproveLoan(payload *dto.ApproveLoanRequest) error {
	loan, err := s.LoanRepository.GetLoanByID(payload.ID)
	if err != nil {
		return err
	}
	if loan.Status != "proposed" {
		return errors.New("Loan status is not eligible for this to proceed")
	}

	src, err := payload.PhotoOfVisit.Open()
	if err != nil {
		return fmt.Errorf("Error when opening photo of visit", err)
	}
	defer src.Close()

	uploadDir := "./uploads/visit-documents"
	os.MkdirAll(uploadDir, os.ModePerm)

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), payload.PhotoOfVisit.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to Create File", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	approvedAt, err := time.Parse("2006-01-02", payload.ApprovalDate)
	if err != nil {
		return fmt.Errorf("invalid approval date format: %w", err)
	}

	data := database.Loan{
		ID:                 payload.ID,
		ApprovalEmployeeID: &payload.EmployeeID,
		AgreementLetter:    filePath, // or SignedAgreementLetter, depending on your use case
		ApprovedAt:         &approvedAt,
		Status:             "approved",
	}

	err = s.LoanRepository.UploadLoanByID(data) // You'll need to implement this
	if err != nil {
		return fmt.Errorf("failed to update loan: %w", err)
	}

	return nil

}
