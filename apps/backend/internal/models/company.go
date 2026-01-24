package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Company represents a tenant/organization
type Company struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Users      []User     `gorm:"foreignKey:CompanyID" json:"-"`
	Vendors    []Vendor   `gorm:"foreignKey:CompanyID" json:"-"`
	Categories []Category `gorm:"foreignKey:CompanyID" json:"-"`
	Bills      []Bill     `gorm:"foreignKey:CompanyID" json:"-"`
}

func (c *Company) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
