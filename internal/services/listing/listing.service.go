package services

import (
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/listing"
)

type IListingService interface {
	CreateListing(payload *model.Listing) error
	GetListingById(listingId string) (*model.Listing, error)
	GetListingsByUserId(userId string) ([]*model.Listing, error)
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

func (s *listingService) GetListingsByUserId(userId string) ([]*model.Listing, error) {
	return s.listingRepo.GetListingsByUserId(userId)
}

func NewListingService(listingRepo repo.IListingRepository) IListingService {
	return &listingService{listingRepo: listingRepo}
}
