package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BillStatus string

const (
	StatusDraft   BillStatus = "draft"
	StatusUnpaid  BillStatus = "unpaid"
	StatusPaid    BillStatus = "paid"
	StatusOverdue BillStatus = "overdue"
)

type RecurringFrequency string

const (
	FrequencyWeekly  RecurringFrequency = "weekly"
	FrequencyMonthly RecurringFrequency = "monthly"
	FrequencyYearly  RecurringFrequency = "yearly"
)

// Bill represents a payable bill
type Bill struct {
	ID                 uuid.UUID           `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyID          uuid.UUID           `gorm:"type:uuid;not null;index" json:"company_id"`
	UserID             uuid.UUID           `gorm:"type:uuid;not null;index" json:"user_id"`
	VendorID           *uuid.UUID          `gorm:"type:uuid;index" json:"vendor_id"`
	CategoryID         *uuid.UUID          `gorm:"type:uuid;index" json:"category_id"`
	Title              string              `gorm:"type:varchar(255);not null" json:"title"`
	InvoiceNumber      *string             `gorm:"type:varchar(100);uniqueIndex" json:"invoice_number"`
	Amount             decimal.Decimal     `gorm:"type:decimal(15,2);not null" json:"amount"`
	Currency           string              `gorm:"type:varchar(10);default:'USD'" json:"currency"`
	DueDate            time.Time           `gorm:"type:date;not null" json:"due_date"`
	PaidDate           *time.Time          `gorm:"type:date" json:"paid_date"`
	Status             BillStatus          `gorm:"type:varchar(20);default:'draft'" json:"status"`
	IsRecurring        bool                `gorm:"default:false" json:"is_recurring"`
	RecurringFrequency *RecurringFrequency `gorm:"type:varchar(20)" json:"recurring_frequency"`
	RecurringDay       *int                `json:"recurring_day"`
	PaymentMethod      *string             `gorm:"type:varchar(100)" json:"payment_method"`
	Notes              *string             `gorm:"type:text" json:"notes"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	DeletedAt          gorm.DeletedAt      `gorm:"index" json:"-"`

	// Relations
	Company     Company          `gorm:"foreignKey:CompanyID" json:"-"`
	User        User             `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Vendor      *Vendor          `gorm:"foreignKey:VendorID" json:"vendor,omitempty"`
	Category    *Category        `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Attachments []BillAttachment `gorm:"foreignKey:BillID" json:"attachments,omitempty"`
	Activities  []BillActivity   `gorm:"foreignKey:BillID" json:"activities,omitempty"`
}

func (b *Bill) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

// IsOverdue checks if the bill is overdue
func (b *Bill) IsOverdue() bool {
	if b.Status == StatusPaid {
		return false
	}
	return time.Now().After(b.DueDate)
}

// DaysUntilDue returns days until due date (negative if overdue)
func (b *Bill) DaysUntilDue() int {
	duration := time.Until(b.DueDate)
	return int(duration.Hours() / 24)
}
