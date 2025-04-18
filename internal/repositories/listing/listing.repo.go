package repo

import (
	"ecommerce/internal/filters"
	"ecommerce/internal/model"
	"errors"

	"gorm.io/gorm"
)

type IListingRepository interface {
	CreateListing(payload *model.Listing) error
	GetListingById(listingId string) (*model.Listing, error)
	GetListingsBySellerId(sellerId string, page int, limit int) ([]*model.Listing, int64, error)
	GetAllListings(page int, limit int, filters *filters.ListingFilter) ([]*model.Listing, int64, error)
	UpdateListing(listingId string, payload *model.Listing) error
	DeleteListing(listingId string) error
}

type listingRepository struct {
	db *gorm.DB
}

func (s *listingRepository) CreateListing(payload *model.Listing) error {
	return s.db.Create(&payload).Error
}

func (s *listingRepository) UpdateListing(listingId string, payload *model.Listing) error {
	return s.db.Model(&model.Listing{}).Where("id = ?", listingId).Updates(payload).Error
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

func (s *listingRepository) GetListingsBySellerId(sellerId string, page int, limit int) ([]*model.Listing, int64, error) {
	var listings []*model.Listing
	var total int64
	err := s.db.Where("seller_id = ?", sellerId).Offset((page - 1) * limit).Limit(limit).Find(&listings).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return listings, total, nil
}

func (s *listingRepository) GetAllListings(page int, limit int, filters *filters.ListingFilter) ([]*model.Listing, int64, error) {
	var listings []*model.Listing
	var total int64

	query := s.db

	if filters.Title != nil {
		query = query.Where("title LIKE ?", "%"+*filters.Title+"%")
	}
	if filters.Price != nil {
		query = query.Where("price = ?", *filters.Price)
	}
	if filters.Area != nil {
		query = query.Where("area = ?", *filters.Area)
	}
	if filters.Bedroom != nil {
		query = query.Where("bedroom = ?", *filters.Bedroom)
	}
	if filters.Bathroom != nil {
		query = query.Where("bathroom = ?", *filters.Bathroom)
	}
	if filters.Floor != nil {
		query = query.Where("floor = ?", *filters.Floor)
	}
	if filters.Direction != nil {
		query = query.Where("direction = ?", *filters.Direction)
	}
	if filters.LegalStatus != nil {
		query = query.Where("legal_status = ?", *filters.LegalStatus)
	}
	if filters.IsForRent != nil {
		query = query.Where("is_for_rent = ?", *filters.IsForRent)
	}
	if filters.IsPublished != nil {
		query = query.Where("is_published = ?", *filters.IsPublished)
	}
	if filters.StartDate != nil && filters.EndDate != nil {
		query = query.Where("start_date = ?", *filters.StartDate).Where("end_date = ?", *filters.EndDate)
	}

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&listings).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return listings, total, nil
}

func (s *listingRepository) DeleteListing(listingId string) error {
	return s.db.Delete(&model.Listing{}, "id = ?", listingId).Error
}

func NewListingRepository(db *gorm.DB) IListingRepository {
	return &listingRepository{db: db}
}
