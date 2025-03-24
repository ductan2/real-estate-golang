package repo

import (
	"ecommerce/internal/model"
	"errors"

	"gorm.io/gorm"
)

type IListingRepository interface {
	CreateListing(payload *model.Listing) error
	GetListingById(listingId string) (*model.Listing, error)
	GetListingsByUserId(userId string) ([]*model.Listing, error)
}

type listingRepository struct {
	db *gorm.DB
}

func (s *listingRepository) CreateListing(payload *model.Listing) error {
	return s.db.Create(&payload).Error
}

func (s *listingRepository) GetListingById(listingId string) (*model.Listing, error) {
	var listing model.Listing
	err := s.db.First(&listing, "id = ?", listingId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &listing, nil
}

func (s *listingRepository) GetListingsByUserId(userId string) ([]*model.Listing, error) {
	var listings []*model.Listing
	err := s.db.Where("user_id = ?", userId).Find(&listings).Error
	if err != nil {
		return nil, err
	}
	return listings, nil
}

func NewListingRepository(db *gorm.DB) IListingRepository {
	return &listingRepository{db: db}
}
