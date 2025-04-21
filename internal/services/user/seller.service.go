package services

import (
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/user"

	"github.com/google/uuid"
)

type ISellerService interface {
	ApplyForSeller(userId string) error
	GetSeller(sellerId string) *model.Seller
	GetAllSeller() ([]*model.Seller, error)
}

type sellerService struct {
	sellerRepo repo.ISellerRepository
}

// GetSeller implements ISellerService.
func (s *sellerService) GetSeller(sellerId string) *model.Seller {
	seller, err := s.sellerRepo.GetSellerById(sellerId)
	if err != nil {
		return nil
	}
	return seller
}

// GetAllSeller implements ISellerService.
func (s *sellerService) GetAllSeller() ([]*model.Seller, error) {
	return s.sellerRepo.GetAllSeller()
}

// ApplyForSeller implements ISellerService.
func (s *sellerService) ApplyForSeller(userId string) error {
	id, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	err = s.sellerRepo.CreateSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func NewSellerService(sellerRepo repo.ISellerRepository) ISellerService {
	return &sellerService{sellerRepo: sellerRepo}
}
