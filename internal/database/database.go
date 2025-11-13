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
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

type Employee struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeNumber string         `gorm:"unique;not null;size:50" json:"employee_number"`
	Name           string         `gorm:"not null;size:255" json:"name"`
	PhoneNumber    string         `gorm:"size:20" json:"phone_number"`
	Email          string         `gorm:"unique;size:255" json:"email"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
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

	// Set connection pool settings
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
		// Add other models here as you create them
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
