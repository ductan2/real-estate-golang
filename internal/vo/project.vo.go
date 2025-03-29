package vo

import "time"

type ProjectCreateRequest struct {
	Name            string     `json:"name" binding:"required"`
	Description     string     `json:"description" binding:"required"`
	LongDescription string     `json:"longDescription" binding:"required"`
	Status          string     `json:"status" binding:"required"`
	AreaLand        float64    `json:"area" binding:"required,gt=0"`
	ProjectType     string     `json:"projectType" binding:"required"`
	Images          []any      `json:"images" binding:"required"`
	LegalStatus     string     `json:"legalStatus" binding:"required"`
	IsPublish       bool       `json:"isPublish"`
	
	Address         string     `json:"address" binding:"required"`
	StartDate       *time.Time `json:"startDate" binding:"required"`
	EndDate         *time.Time `json:"endDate" binding:"required"`
	Apartment       *int       `json:"areaBuild"`
	InvestorID      string     `json:"investorId" binding:"required,uuid"`
}

type ProjectUpdateRequest struct {
	Name            *string   `json:"name"`
	Description     *string   `json:"description"`
	LongDescription *string   `json:"longDescription"`
	Status          *string   `json:"status"`
	AreaLand        *float64  `json:"areaLand"`
	AreaBuild       *float64  `json:"areaBuild"`
	ProjectType     *string   `json:"projectType"`
	Images          *[]string `json:"images"`
	Address         *string   `json:"address"`
	Apartment       *int      `json:"apartment"`
}
