package vo

import "github.com/google/uuid"

type Price struct {
	MinPrice float64 `json:"min_price" binding:"required"`
	MaxPrice float64 `json:"max_price" binding:"required"`
	Unit     string  `json:"unit" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type ProjectCreateRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	LongDescription string    `json:"long_description" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	Area        float64   `json:"area" binding:"required"`
	Images      []string  `json:"images" binding:"required"`
	Apartment   *int      `json:"apartment"`
	AreaBuild   *float64  `json:"area_build"`
	ProjectType string    `json:"project_type" binding:"required"`
	Price       Price     `json:"price" binding:"required"`
	InvestorID  uuid.UUID `json:"investor_id" binding:"required"`
}
