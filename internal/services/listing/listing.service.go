package services

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/listing"
)

type IListingService interface {
	CreateListing(payload *model.Listing) error
	GetListingById(listingId string) (*model.Listing, error)
	GetListingsBySellerId(sellerId string, page int, limit int) ([]*model.Listing, int64, error)
	GetAllListings(page int, limit int, filters *filters.ListingFilter) ([]*model.Listing, int64, error)
	UpdateListing(listingId string, payload *model.Listing) error
	DeleteListing(listingId string) error
}

type listingService struct {
	listingRepo repo.IListingRepository
}

func (s *listingService) CreateListing(payload *model.Listing) error {
	return s.listingRepo.CreateListing(payload)
}

func (s *listingService) GetListingById(listingId string) (*model.Listing, error) {
	return s.listingRepo.GetListingById(listingId)
}

func (s *listingService) GetListingsBySellerId(sellerId string, page int, limit int) ([]*model.Listing, int64, error) {
	return s.listingRepo.GetListingsBySellerId(sellerId, page, limit)
}

func (s *listingService) GetAllListings(page int, limit int, filters *filters.ListingFilter) ([]*model.Listing, int64, error) {
	return s.listingRepo.GetAllListings(page, limit, filters)
}

func (s *listingService) UpdateListing(listingId string, payload *model.Listing) error {
	return s.listingRepo.UpdateListing(listingId, payload)
}

func (s *listingService) DeleteListing(listingId string) error {
	return s.listingRepo.DeleteListing(listingId)
}

func NewListingService(listingRepo repo.IListingRepository) IListingService {
	return &listingService{listingRepo: listingRepo}
}
