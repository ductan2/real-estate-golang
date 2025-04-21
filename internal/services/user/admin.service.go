package services

import (
	"ecommerce/global"
	repo "ecommerce/internal/repositories/user"
	mail "ecommerce/internal/utils/email"
	"fmt"

	"github.com/google/uuid"
)

type IAdminService interface {
	ApplyForAdmin(userId string) error
	ApproveSellerRequest(userId, sellerId, sellerEmail string, approved bool) error
	BlockSeller(userId, sellerId string, reason string) error
	CheckAdmin(userId string) (bool, error)
}

type adminService struct {
	adminRepo repo.IAdminRepository
}

// BlockSeller implements IAdminService.
func (s *adminService) BlockSeller(userId string, sellerId string, reason string) error {
	id, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	err = s.adminRepo.BlockSeller(id, sellerId, reason)
	if err != nil {
		return err
	}
	return nil
}

// ApproveSellerRequest implements IAdminService.
func (s *adminService) ApproveSellerRequest(userId string, sellerId, sellerEmail string, approved bool) error {
	fmt.Println("userId", userId)
	fmt.Println("sellerId", sellerId)
	fmt.Println("sellerEmail", sellerEmail)
	fmt.Println("approved", approved)
	id, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	fmt.Println("sellerId", sellerId)
	err = s.adminRepo.ApproveSellerRequest(id, sellerId, approved)
	if err != nil {
		return err
	}
	fmt.Println("approved", approved)

	if approved {
		go func (){
			mail.SendEmail([]string{sellerEmail}, global.Config.SMTP.Username, "Your account has been approved", "Your account has been approved")
		}()
	}
	return nil
}

// ApplyForAdmin implements IAdminService.
func (s *adminService) ApplyForAdmin(userId string) error {
	id, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	err = s.adminRepo.CreateAdmin(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *adminService) CheckAdmin(userId string) (bool, error) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return false, err
	}
	user, err := s.adminRepo.CheckAdmin(id)
	if err != nil {
		return false, err
	}
	return user, nil
}
func NewAdminService(adminRepo repo.IAdminRepository) IAdminService {
	return &adminService{adminRepo: adminRepo}
}
