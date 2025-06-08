package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;index;not null"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	Read      bool      `json:"read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
