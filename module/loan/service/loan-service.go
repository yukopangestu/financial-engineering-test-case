package service

import (
	"errors"
	"financial-engineering-test-case/internal/config"
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

	"github.com/jung-kurt/gofpdf"
	"gopkg.in/gomail.v2"
)

type LoanService struct {
	LoanRepository  *repository.LoanRepository
	BorrowerService *bService.BorrowerService
	Config          *config.Config
}

var _ domain.LoanService = (*LoanService)(nil)

func NewLoanService(
	LoanRepository *repository.LoanRepository,
	BorrowerService *bService.BorrowerService,
	Config *config.Config,
) *LoanService {
	return &LoanService{
		LoanRepository:  LoanRepository,
		BorrowerService: BorrowerService,
		Config:          Config,
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
		Status:     "proposed",
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
		ProofOfVisit:       filePath,
		ApprovedAt:         &approvedAt,
		Status:             "approved",
	}

	err = s.LoanRepository.UploadLoanByID(data)
	if err != nil {
		return fmt.Errorf("failed to update loan: %w", err)
	}

	return nil
}

func (s LoanService) sendInvestorEmail(investorEmail, investorName string, loanNumber string, investmentAmount float64, pdfPath string) error {
	if s.Config.SMTPUsername == "" || s.Config.SMTPPassword == "" {
		return nil
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.Config.SMTPFrom)
	m.SetHeader("To", investorEmail)
	m.SetHeader("Subject", "Investment Confirmation - Loan "+loanNumber)

	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
			<h2 style="color: #2c3e50;">Investment Confirmation</h2>
			<p>Dear %s,</p>
			<p>Thank you for your investment in our loan program. This email confirms your investment details:</p>

			<table style="border-collapse: collapse; margin: 20px 0;">
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd; background-color: #f9f9f9;"><strong>Loan Number:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
				</tr>
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd; background-color: #f9f9f9;"><strong>Investment Amount:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">$%.2f</td>
				</tr>
				<tr>
					<td style="padding: 8px; border: 1px solid #ddd; background-color: #f9f9f9;"><strong>Date:</strong></td>
					<td style="padding: 8px; border: 1px solid #ddd;">%s</td>
				</tr>
			</table>

			<p>Please find the loan investment agreement letter attached to this email for your records.</p>

			<p>The loan will be processed and disbursed according to the terms outlined in the agreement. You will receive further updates on the disbursement of the loan.</p>

			<p style="margin-top: 30px;">Best regards,<br>
			<strong>Loan Management Team</strong></p>

			<hr style="border: none; border-top: 1px solid #ddd; margin: 30px 0;">
			<p style="font-size: 12px; color: #777;">This is an automated message. Please do not reply to this email.</p>
		</body>
		</html>
	`, investorName, loanNumber, investmentAmount, time.Now().Format("January 2, 2006"))

	m.SetBody("text/html", body)

	if pdfPath != "" {
		m.Attach(pdfPath)
	}

	d := gomail.NewDialer(s.Config.SMTPHost, s.Config.SMTPPort, s.Config.SMTPUsername, s.Config.SMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", investorEmail, err)
	}

	return nil
}

func (s LoanService) InvestLoan(payload *dto.InvestLoanRequest, id uint) error {
	loan, err := s.LoanRepository.GetLoanByID(id)
	if err != nil {
		return err
	}
	if loan.Status != "approved" {
		return errors.New("Loan status is not eligible for this to proceed")
	}

	totalInvestment := 0.0
	for _, investor := range payload.Investors {
		totalInvestment += investor.InvestmentAmount
	}

	if totalInvestment != loan.Amount {
		return fmt.Errorf("total investment amount ($%.2f) must equal loan amount ($%.2f)", totalInvestment, loan.Amount)
	}

	var loanInvestors []database.LoanInvestor
	var investorDetails []database.Investor

	for _, investor := range payload.Investors {
		inv, err := s.LoanRepository.GetInvestorByID(investor.ID)
		if err != nil {
			return fmt.Errorf("investor with ID %d not found: %w", investor.ID, err)
		}
		if inv.ID == 0 {
			return fmt.Errorf("investor with ID %d does not exist", investor.ID)
		}

		loanInvestor := database.LoanInvestor{
			LoanID:           id,
			InvestorID:       investor.ID,
			InvestmentAmount: investor.InvestmentAmount,
		}

		err = s.LoanRepository.CreateLoanInvestor(loanInvestor)
		if err != nil {
			return fmt.Errorf("failed to create loan-investor record: %w", err)
		}

		loanInvestors = append(loanInvestors, loanInvestor)
		investorDetails = append(investorDetails, inv)
	}

	pdfPath, err := s.generateAgreementLetterPDF(loan, loanInvestors)
	if err != nil {
		return fmt.Errorf("failed to generate agreement letter: %w", err)
	}

	investedAt := time.Now()
	updatedLoan := database.Loan{
		ID:              id,
		Status:          "invested",
		InvestedAt:      &investedAt,
		AgreementLetter: pdfPath,
	}

	err = s.LoanRepository.UploadLoanByID(updatedLoan)
	if err != nil {
		return fmt.Errorf("failed to update loan: %w", err)
	}

	for i, investor := range loanInvestors {
		inv := investorDetails[i]
		err = s.sendInvestorEmail(inv.Email, inv.Name, loan.LoanNumber, investor.InvestmentAmount, pdfPath)
		if err != nil {
			fmt.Printf("Warning: failed to send email to investor %s: %v\n", inv.Email, err)
		}
	}

	return nil
}

func (s LoanService) generateAgreementLetterPDF(loan database.Loan, investors []database.LoanInvestor) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "LOAN INVESTMENT AGREEMENT LETTER")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Date: %s", time.Now().Format("January 2, 2006")))
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Loan Details")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 11)
	pdf.Cell(50, 8, "Loan Number:")
	pdf.Cell(0, 8, loan.LoanNumber)
	pdf.Ln(6)

	pdf.Cell(50, 8, "Loan Amount:")
	pdf.Cell(0, 8, fmt.Sprintf("$%.2f", loan.Amount))
	pdf.Ln(6)

	pdf.Cell(50, 8, "Interest Rate:")
	pdf.Cell(0, 8, fmt.Sprintf("%.2f%%", loan.Interest))
	pdf.Ln(6)

	pdf.Cell(50, 8, "Borrower ID:")
	pdf.Cell(0, 8, fmt.Sprintf("%d", loan.BorrowerID))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Investment Details")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 11)
	totalInvestment := 0.0

	for i, inv := range investors {
		pdf.SetFont("Arial", "B", 11)
		pdf.Cell(0, 8, fmt.Sprintf("Investor %d:", i+1))
		pdf.Ln(6)

		pdf.SetFont("Arial", "", 11)
		pdf.Cell(10, 6, "")
		pdf.Cell(50, 6, "Investor ID:")
		pdf.Cell(0, 6, fmt.Sprintf("%d", inv.InvestorID))
		pdf.Ln(6)

		pdf.Cell(10, 6, "")
		pdf.Cell(50, 6, "Investment Amount:")
		pdf.Cell(0, 6, fmt.Sprintf("$%.2f", inv.InvestmentAmount))
		pdf.Ln(6)

		totalInvestment += inv.InvestmentAmount
	}

	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(60, 8, "Total Investment:")
	pdf.Cell(0, 8, fmt.Sprintf("$%.2f", totalInvestment))
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "Terms and Conditions")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 10)
	terms := []string{
		"1. The investors agree to invest the specified amounts in the loan.",
		"2. The loan will be disbursed to the borrower after all investment funds are received.",
		"3. The borrower agrees to repay the loan with the specified interest rate.",
		"4. Returns will be distributed to investors proportionally to their investment amounts.",
		"5. This agreement is binding and governed by applicable laws.",
	}

	for _, term := range terms {
		pdf.MultiCell(0, 6, term, "", "", false)
		pdf.Ln(2)
	}

	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Signatures")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(90, 20, "________________________")
	pdf.Cell(90, 20, "________________________")
	pdf.Ln(6)
	pdf.Cell(90, 6, "Authorized Officer")
	pdf.Cell(90, 6, "Date")

	uploadDir := "./uploads/agreement-letters"
	os.MkdirAll(uploadDir, os.ModePerm)

	fileName := fmt.Sprintf("agreement_%s_%d.pdf", loan.LoanNumber, time.Now().Unix())
	fileName = filepath.Join(uploadDir, filepath.Base(fileName))

	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	return fileName, nil
}

func (s LoanService) DisburseLoan(payload *dto.DisbursedLoanRequest, id uint) error {
	loan, err := s.LoanRepository.GetLoanByID(id)
	if err != nil {
		return fmt.Errorf("failed to get loan: %w", err)
	}

	if loan.Status != "invested" {
		return errors.New("loan status is not eligible for disbursement (must be 'invested')")
	}

	src, err := payload.SignedAgreementLetter.Open()
	if err != nil {
		return fmt.Errorf("error when opening signed agreement letter: %w", err)
	}
	defer src.Close()

	uploadDir := "./uploads/signed-agreement-letters"
	os.MkdirAll(uploadDir, os.ModePerm)

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), payload.SignedAgreementLetter.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to save signed agreement letter: %w", err)
	}

	disbursementDate, err := time.Parse("2006-01-02", payload.DisbursementDate)
	if err != nil {
		return fmt.Errorf("invalid disbursement date format (expected YYYY-MM-DD): %w", err)
	}

	updatedLoan := database.Loan{
		ID:                    id,
		DisbursedEmployeeID:   &payload.EmployeeID,
		SignedAgreementLetter: filePath,
		DisbursementDate:      &disbursementDate,
		Status:                "disbursed",
	}

	err = s.LoanRepository.UploadLoanByID(updatedLoan)
	if err != nil {
		return fmt.Errorf("failed to update loan with disbursement info: %w", err)
	}

	return nil
}
