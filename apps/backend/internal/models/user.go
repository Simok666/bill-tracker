package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
)

// User represents a user account
type User struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"company_id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Email         string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash  string         `gorm:"type:varchar(255)" json:"-"`
	Role          UserRole       `gorm:"type:varchar(50);default:'member'" json:"role"`
	AvatarURL     *string        `gorm:"type:text" json:"avatar_url"`
	EmailVerified bool           `gorm:"default:false" json:"email_verified"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Company    Company       `gorm:"foreignKey:CompanyID" json:"-"`
	Sessions   []Session     `gorm:"foreignKey:UserID" json:"-"`
	Bills      []Bill        `gorm:"foreignKey:UserID" json:"-"`
	Activities []BillActivity `gorm:"foreignKey:UserID" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
