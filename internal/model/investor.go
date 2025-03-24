package model

import (
	"time"

	"github.com/google/uuid"
)

type Investor struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Name        string    `json:"name" gorm:"type:varchar(100);not null"`
    Address     string    `json:"address" gorm:"type:varchar(200);not null"`
    Email       string    `json:"email" gorm:"type:varchar(100);not null"`
    Phone       string    `json:"phone" gorm:"type:varchar(50);notnull"`
    Website     string    `json:"website" gorm:"type:varchar(255);null"`
    Description string    `json:"description" gorm:"type:text;null"`
    Logo        string    `json:"logo" gorm:"type:varchar(255);null"`
	Background  string    `json:"background" gorm:"type:varchar(255);null"`
	Type        string    `json:"type" gorm:"type:varchar(50);not null"`

    UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
    User        User      `json:"user" gorm:"foreignKey:UserId"`
    CreatedAt   time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

func (Investor) TableName() string {
	return "investors"
}
