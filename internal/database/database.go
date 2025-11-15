package database

import (
	"financial-engineering-test-case/internal/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Borrower struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	BorrowerNumber string         `gorm:"unique;not null;size:50" json:"borrower_number"`
	Name           string         `gorm:"not null;size:255" json:"name"`
	PhoneNumber    string         `gorm:"size:20" json:"phone_number"`
	Email          string         `gorm:"unique;size:255" json:"email"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Employee struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeNumber string         `gorm:"unique;not null;size:50" json:"employee_number"`
	Name           string         `gorm:"not null;size:255" json:"name"`
	PhoneNumber    string         `gorm:"size:20" json:"phone_number"`
	Email          string         `gorm:"unique;size:255" json:"email"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Loan struct {
	ID                    uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanNumber            string         `gorm:"size:255" json:"loan_number"`
	ApprovalEmployeeID    *uint          `gorm:"index" json:"approval_employee_id,omitempty"`
	DisbursedEmployeeID   *uint          `gorm:"index" json:"disbursed_employee_id,omitempty"`
	BorrowerID            uint           `gorm:"not null;index" json:"borrower_id"`
	Status                string         `gorm:"size:50;check:status IN ('proposed','approved','invested','disbursed')" json:"status"`
	Amount                float64        `gorm:"type:decimal(10,2)" json:"amount"`
	Interest              float64        `gorm:"type:decimal(10,2)" json:"interest"`
	ProofOfVisit          string         `gorm:"text" json:"proof_of_visit"`
	AgreementLetter       string         `gorm:"text" json:"agreement_letter"`
	SignedAgreementLetter string         `gorm:"text" json:"signed_agreement_letter"`
	ApprovedAt            *time.Time     `gorm:"type:datetime" json:"approved_at,omitempty"`
	InvestedAt            *time.Time     `gorm:"type:datetime" json:"invested_at,omitempty"`
	DisbursementDate      *time.Time     `gorm:"type:datetime" json:"disbursement_date,omitempty"`
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Borrower          Borrower  `gorm:"foreignKey:BorrowerID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"borrower,omitempty"`
	ApprovalEmployee  *Employee `gorm:"foreignKey:ApprovalEmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"approval_employee,omitempty"`
	DisbursedEmployee *Employee `gorm:"foreignKey:DisbursedEmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"disbursed_employee,omitempty"`
}

type Investor struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	InvestorNumber string         `gorm:"size:255" json:"investor_number"`
	Name           string         `gorm:"not null;size:255" json:"name"`
	PhoneNumber    string         `gorm:"size:20" json:"phone_number"`
	Email          string         `gorm:"unique;size:255" json:"email"`
	InvestedAmount float64        `gorm:"type:decimal(10,2)" json:"invested_amount"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type LoanInvestor struct {
	ID               uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID           uint           `gorm:"not null;index" json:"loan_id"`
	InvestorID       uint           `gorm:"not null;index" json:"investor_id"`
	InvestmentAmount float64        `gorm:"type:decimal(10,2)" json:"investment_amount"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	Loan     Loan     `gorm:"foreignKey:LoanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"loan,omitempty"`
	Investor Investor `gorm:"foreignKey:InvestorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"investor,omitempty"`
}

var DB *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established successfully")
	return DB, nil
}

func AutoMigrate() error {
	log.Println("Running auto-migration...")

	err := DB.AutoMigrate(
		Borrower{},
		Employee{},
		Investor{},
		Loan{},
		LoanInvestor{},
	)

	if err != nil {
		return err
	}

	log.Println("Auto-migration completed successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
