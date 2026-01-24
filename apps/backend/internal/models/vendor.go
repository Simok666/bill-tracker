package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Vendor represents a bill vendor/supplier
type Vendor struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"company_id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	LogoURL      *string        `gorm:"type:text" json:"logo_url"`
	ContactEmail *string        `gorm:"type:varchar(255)" json:"contact_email"`
	Website      *string        `gorm:"type:varchar(255)" json:"website"`
	ContactInfo  *string        `gorm:"type:text" json:"contact_info"`
	Address      *string        `gorm:"type:text" json:"address"`
	Location     *string        `gorm:"type:varchar(255)" json:"location"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Company Company `gorm:"foreignKey:CompanyID" json:"-"`
	Bills   []Bill  `gorm:"foreignKey:VendorID" json:"-"`
}

func (v *Vendor) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}
