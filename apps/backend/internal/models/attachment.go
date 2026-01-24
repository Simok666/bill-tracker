package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BillAttachment represents an invoice or document attachment
type BillAttachment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BillID     uuid.UUID `gorm:"type:uuid;not null;index" json:"bill_id"`
	Filename   string    `gorm:"type:varchar(255);not null" json:"filename"`
	FileURL    string    `gorm:"type:text;not null" json:"file_url"`
	FileType   string    `gorm:"type:varchar(50);not null" json:"file_type"`
	FileSize   int64     `gorm:"not null" json:"file_size"`
	UploadedAt time.Time `json:"uploaded_at"`

	// Relations
	Bill Bill `gorm:"foreignKey:BillID" json:"-"`
}

func (a *BillAttachment) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	if a.UploadedAt.IsZero() {
		a.UploadedAt = time.Now()
	}
	return nil
}
