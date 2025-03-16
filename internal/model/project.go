package model

import (
	"time"

	"github.com/google/uuid"
)

type Price struct {
	MinPrice float64 `json:"min_price" gorm:"type:decimal(10,2);not null"`
	MaxPrice float64 `json:"max_price" gorm:"type:decimal(10,2);not null"`
	Unit     string  `json:"unit" gorm:"type:varchar(30);not null"`
	Currency string  `json:"currency" gorm:"type:varchar(30);not null"`
}

type Location struct {
	Latitude  float64 `json:"latitude" gorm:"type:decimal(10,2);not null"`
	Longitude float64 `json:"longitude" gorm:"type:decimal(10,2);not null"`
	Address   string  `json:"address" gorm:"type:varchar(255);not null"`
}

type Project struct {
	ID              uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name            string         `json:"name" gorm:"type:varchar(100);not null"`
	Description     string         `json:"description" gorm:"type:varchar(255);not null"`
	LongDescription string         `json:"long_description" gorm:"type:text;not null"`
	Status          string         `json:"status" gorm:"type:varchar(30);not null"`
	AreaLand        float64        `json:"area_land" gorm:"type:decimal(10,2);not null"`
	AreaBuild       float64        `json:"area_build" gorm:"type:decimal(10,2);not null"`

	ProjectType     string         `json:"project_type" gorm:"type:varchar(30);not null"`
	Price           Price          `json:"price" gorm:"type:jsonb;not null"`
	Images          []string       `json:"images" gorm:"type:jsonb;not null"`
	Video           string         `json:"video" gorm:"type:varchar(255);null"`
	Location        Location       `json:"location" gorm:"type:jsonb;not null"`

	InvestorID      uuid.UUID      `json:"investor_id" gorm:"type:uuid;not null"`
	Investor        Investor       `json:"investor" gorm:"foreignKey:InvestorID"`

	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type LoanSupport struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProjectID     uuid.UUID `json:"project_id" gorm:"type:uuid;not null"`
	LoanRate      float64   `json:"loan_rate" gorm:"type:decimal(10,2);not null"`
	MinLoanAmount float64   `json:"min_loan_amount" gorm:"type:decimal(10,2);not null"`
	MaxLoanAmount float64   `json:"max_loan_amount" gorm:"type:decimal(10,2);not null"`
	MinLoanPeriod int       `json:"min_loan_period" gorm:"type:int;not null"`
	MaxLoanPeriod int       `json:"max_loan_period" gorm:"type:int;not null"`
	LoanType      string    `json:"loan_type" gorm:"type:varchar(30);not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProjectManager struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProjectID uuid.UUID `json:"project_id" gorm:"type:uuid;not null"`
	Project   Project   `json:"project" gorm:"foreignKey:ProjectID"`
	
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`

	Role      string    `json:"role" gorm:"type:varchar(30);not null"` // admin, manager, viewer
	Permission string    `json:"permission" gorm:"type:jsonb;not null"` // view only, full access, edit only, delete only

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Project) TableName() string {
	return "projects"
}

func (LoanSupport) TableName() string {
	return "loan_supports"
}

func (ProjectManager) TableName() string {
	return "project_managers"
}
