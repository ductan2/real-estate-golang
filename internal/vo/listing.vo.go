package vo

import "time"

type ListingCreateRequest struct {
	Title           string    `json:"title" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	LongDescription string    `json:"long_description" binding:"required"`
	Price           float64   `json:"price" binding:"required"`
	Unit            string    `json:"unit" binding:"required"`
	PropertyType    string    `json:"property_type" binding:"required"`
	Area            float64   `json:"area" binding:"required"`
	Bedroom         int       `json:"bedroom" binding:"required"`
	Bathroom        int       `json:"bathroom" binding:"required"`
	Floor           int       `json:"floor" binding:"required"`
	Direction       string    `json:"direction" binding:"required"`
	LegalStatus     string    `json:"legal_status" binding:"required"`
	IsForRent       bool      `json:"is_for_rent" binding:"required"`
	VideoURL        string    `json:"video_url" binding:"required"`
	DurationListing int       `json:"duration_listing" binding:"required"`
	Images          []string  `json:"images" binding:"required"`
	Category        string    `json:"category" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	IsPublished     bool      `json:"is_published" binding:"required"`
	StartDate       time.Time `json:"start_date" binding:"required"`
	EndDate         time.Time `json:"end_date" binding:"required"`
}

type ListingUpdateRequest struct {
	Title           string    `json:"title" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Price           float64   `json:"price" binding:"required"`
	Images          []string  `json:"images" binding:"required"`
	Unit            string    `json:"unit" binding:"required"`
	PropertyType    string    `json:"property_type" binding:"required"`
	Area            float64   `json:"area" binding:"required"`
	Bedroom         int       `json:"bedroom" binding:"required"`
	Bathroom        int       `json:"bathroom" binding:"required"`
	Floor           int       `json:"floor" binding:"required"`
	Direction       string    `json:"direction" binding:"required"`
	LegalStatus     string    `json:"legal_status" binding:"required"`
	IsForRent       bool      `json:"is_for_rent" binding:"required"`
	VideoURL        string    `json:"video_url" binding:"required"`
	DurationListing int       `json:"duration_listing" binding:"required"`
	IsPublished     bool      `json:"is_published" binding:"required"`
	StartDate       time.Time `json:"start_date" binding:"required"`
	EndDate         time.Time `json:"end_date" binding:"required"`
}
