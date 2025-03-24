package model

import (
	"time"

	"github.com/google/uuid"
)

type SubProject struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProjectID uuid.UUID `json:"project_id" gorm:"type:uuid;not null"`
	Project   Project   `json:"project" gorm:"foreignKey:ProjectID"`

	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	LongDescription string    `json:"long_description" gorm:"type:text;not null"`
	Area      float64   `json:"area" gorm:"type:decimal(10,2);not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Floor     int       `json:"floor" gorm:"type:int;null"`
	Images    []string  `json:"images" gorm:"type:jsonb;not null"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Listing struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SubProjectID uuid.UUID `json:"sub_project_id" gorm:"type:uuid;not null"`
	SubProject SubProject `json:"sub_project" gorm:"foreignKey:SubProjectID"`
	Title           string    `json:"title"`
    Price           float64   `json:"price"`
	Unit            string    `json:"unit"` // Unit of the listing
	PropertyType    string    `json:"property_type"` // Type of the listing
    Area            float64   `json:"area"`       // Area of the listing
    Bedroom         int       `json:"bedroom"`    // Number of bedrooms
    Bathroom        int       `json:"bathroom"`   // Number of bathrooms
    Floor           int       `json:"floor"`      // Floor number in the building
    Direction       string    `json:"direction"`  // Direction of the listing
    IsForRent       bool      `json:"is_for_rent"`// true = for rent, false = for sale
    FurnitureStatus string    `json:"furniture_status"` // Furniture status of the listing
    LegalStatus     string    `json:"legal_status"` // Legal status of the listing
    Images          []string  `gorm:"type:jsonb" json:"images"`
	VideoURL        *string    `json:"video_url" gorm:"default:null"` // Video URL of the listing
    Description     string    `gorm:"type:text" json:"description"`
	LongDescription string    `gorm:"type:text" json:"long_description"`
	IsPublished     bool      `json:"is_published" gorm:"default:false"` // true = published, false = not published
	StartDate       time.Time `json:"start_date"` // Start date of the listing
	EndDate         time.Time `json:"end_date"` // End date of the listing
	ListingType     string    `json:"listing_type"` // Type of the listing [diamond, gold, silver, normal]
	DurationListing int       `json:"duration_listing"` // Duration of the listing in days
	
	SellerID        uuid.UUID `json:"seller_id" gorm:"type:uuid;not null"`
	Seller          Seller    `json:"seller" gorm:"foreignKey:SellerID"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (SubProject) TableName() string {
	return "sub_projects"
}

func (Listing) TableName() string {
	return "listings"
}
