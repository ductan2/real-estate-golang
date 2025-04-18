package filters

import "time"

type ProjectFilter struct {
	Name        *string
	Status      *string
	ProjectType *string
	LegalStatus *string
	MaxArea     *string
	MinArea     *string
	Province    *string
	IsPublish   *bool
	InvestorID  *string
}

type InvestorFilter struct {
	Email   *string
	Phone   *string
	Address *string
	Website *string
}

type ListingFilter struct {
	Title       *string
	Price       *string
	Area        *string
	Bedroom     *string
	Bathroom    *string
	Floor       *string
	Direction   *string
	LegalStatus *string
	IsForRent   *bool
	IsPublished *bool
	StartDate   *time.Time
	EndDate     *time.Time
}
