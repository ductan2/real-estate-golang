package repo

import (
	"ecommerce/internal/model"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISellerRepository interface {
	CreateSeller(userId uuid.UUID) error
	GetSellerById(sellerId string) (*model.Seller, error)
}

type sellerRepository struct {
	db *gorm.DB
}

// CreateSeller implements ISellerRepository.
func (s *sellerRepository) CreateSeller(userId uuid.UUID) error {
	return s.db.Create(&model.Seller{
		UserId: userId,
	}).Error
}

// GetSellerById implements ISellerRepository.
func (s *sellerRepository) GetSellerById(sellerId string) (*model.Seller, error) {
	var seller model.Seller
	err := s.db.Preload("User").First(&seller, "id = ?", sellerId).Error
	if err != nil {
		fmt.Println("err", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &seller, nil
}

func NewSellerRepository(db *gorm.DB) ISellerRepository {
	return &sellerRepository{db: db}
}
