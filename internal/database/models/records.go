package models

import (
	"time"

	"github.com/google/uuid"
)

type RecordType string
type RecordCategory string

const (
	IncomeRecordType RecordType = "income"
	SpendRecordType  RecordType = "spend"
)

const (
	Salary        RecordCategory = "salary"
	Freelance     RecordCategory = "freelance"
	Investment    RecordCategory = "investment"
	Bonus         RecordCategory = "bonus"
	Rental        RecordCategory = "rental_income"
	OtherIncome   RecordCategory = "other_income"
	Food          RecordCategory = "food"
	Transport     RecordCategory = "transport"
	Utilities     RecordCategory = "utilities"
	Rent          RecordCategory = "rent"
	Healthcare    RecordCategory = "healthcare"
	Shopping      RecordCategory = "shopping"
	Education     RecordCategory = "education"
	Entertainment RecordCategory = "entertainment"
	Insurance     RecordCategory = "insurance"
	OtherExpense  RecordCategory = "other_expense"
)

type Record struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Amount      float64        `gorm:"not null"`
	UserId      uuid.UUID      `gorm:"type:uuid;not null"`
	user        User           `gorm:"foreignKey:UserId"` // foreign key for the user who created the entry
	Type        RecordType     `gorm:"not null"`
	Category    RecordCategory `gorm:"not null"`
	Description string
	Date        time.Time `gorm:"not null"`
	IsDeleted   bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time
}
