package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivityAction string

const (
	ActionCreated             ActivityAction = "created"
	ActionUpdated             ActivityAction = "updated"
	ActionStatusChanged       ActivityAction = "status_changed"
	ActionPaymentReminderSent ActivityAction = "payment_reminder_sent"
	ActionAttachmentAdded     ActivityAction = "attachment_added"
	ActionAttachmentRemoved   ActivityAction = "attachment_removed"
	ActionDeleted             ActivityAction = "deleted"
)

// BillActivity represents an activity log entry for a bill
type BillActivity struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BillID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"bill_id"`
	UserID    *uuid.UUID     `gorm:"type:uuid;index" json:"user_id"`
	Action    ActivityAction `gorm:"type:varchar(100);not null" json:"action"`
	Details   *string        `gorm:"type:text" json:"details"`
	CreatedAt time.Time      `json:"created_at"`

	// Relations
	Bill Bill  `gorm:"foreignKey:BillID" json:"-"`
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (a *BillActivity) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
