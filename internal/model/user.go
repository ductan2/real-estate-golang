package model

import (
	"time"

	"ecommerce/pkg/enum"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username       string    `json:"username" gorm:"uniqueIndex;not null"`
	Email          string    `json:"email" gorm:"uniqueIndex;not null"`
	Password       string    `json:"-" gorm:"not null"` // "-" means this field won't be included in JSON
	Role           enum.Role `json:"role" gorm:"type:varchar(20);not null;default:'user'"`
	UserLoginTime  time.Time `json:"-" gorm:"default:null"`
	UserLoginIP    string    `json:"-" gorm:"default:null"`

	UserLogoutTime time.Time `json:"-" gorm:"default:null"`
	UserSalt       string    `json:"-" gorm:"default:null"` // user_salt is the salt that will be used to hash the password
	Verified       bool      `json:"verified" gorm:"default:false"`
	VerifiedAt     time.Time `json:"-" gorm:"default:null"`
	UserInfo       *UserInfo `json:"user_info" gorm:"foreignKey:UserId"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserSession struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	IpAddress string    `json:"ip_address" gorm:"default:null"`
	Location  string    `json:"location" gorm:"default:null"`
	Device    string    `json:"device" gorm:"default:null"`
	UserAgent string    `json:"user_agent" gorm:"default:null"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;constraint:OnDelete:CASCADE"`
	User      *User     `json:"user" gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfo struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId uuid.UUID `json:"user_id" gorm:"type:uuid;not null;constraint:OnDelete:CASCADE"`

	Phone           *string    `json:"phone" gorm:"default:null"`
	Province        *string    `json:"province" gorm:"default:null"`
	Ward            *string    `json:"ward" gorm:"default:null"`
	District        *string    `json:"district" gorm:"default:null"`
	Address         *string    `json:"address" gorm:"default:null"`
	BirthDate       *time.Time `json:"birth_date" gorm:"default:null"`
	Avatar          *string    `json:"avatar" gorm:"default:null"`
	Bio             *string    `json:"bio" gorm:"default:null"`
	Gender          *string    `json:"gender" gorm:"type:varchar(10);default:null"`
	PersonalTaxCode *string    `json:"personal_tax_code" gorm:"default:null"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type Seller struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserId         uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;constraint:OnDelete:CASCADE"`
	IsVerified     bool       `json:"is_verified" gorm:"default:null"`
	VerifiedAt     *time.Time `json:"verified_at" gorm:"default:null"`
	VerifiedBy     *uuid.UUID `json:"verified_by" gorm:"type:uuid;default:null"`
	VerifiedByUser *User      `json:"verified_by_user" gorm:"foreignKey:VerifiedBy"`

	BlockedAt     *time.Time `json:"blocked_at" gorm:"default:null"`
	BlockedBy     *uuid.UUID `json:"blocked_by" gorm:"type:uuid;default:null"`
	BlockedByUser *User      `json:"blocked_by_user" gorm:"foreignKey:BlockedBy"`
	BlockedReason *string    `json:"blocked_reason" gorm:"default:null"`
	User          *User      `json:"user" gorm:"foreignKey:UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

func (UserInfo) TableName() string {
	return "user_infos"
}

func (Seller) TableName() string {
	return "sellers"
}

func (UserSession) TableName() string {
	return "user_sessions"
}

// AfterCreate is a hook that runs after creating a new user
func (u *User) AfterCreate(tx *gorm.DB) error {
	userInfo := UserInfo{
		UserId: u.ID,
	}
	tx.Create(&userInfo)
	return nil
}
