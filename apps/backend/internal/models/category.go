package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Category represents a bill category
type Category struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"company_id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description *string        `gorm:"type:text" json:"description"`
	Icon        *string        `gorm:"type:varchar(50)" json:"icon"`
	Color       *string        `gorm:"type:varchar(20)" json:"color"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Company Company `gorm:"foreignKey:CompanyID" json:"-"`
	Bills   []Bill  `gorm:"foreignKey:CategoryID" json:"-"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
