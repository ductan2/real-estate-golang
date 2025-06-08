package repo

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/enum"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAdminRepository interface {
	CreateAdmin(userId uuid.UUID) error
	ApproveSellerRequest(userId uuid.UUID, sellerId string, approved bool) error
	BlockSeller(userId uuid.UUID, sellerId string, reason string) error
	CheckAdmin(userId uuid.UUID) (bool, error)
	GetAllAdmins() ([]*model.User, error)
}

type adminRepository struct {
	db *gorm.DB
}

// CreateAdmin implements IAdminRepository.
func (a *adminRepository) CreateAdmin(userId uuid.UUID) error {
	return a.db.Model(&model.User{}).Where("id = ?", userId).Update("role", enum.UserRole.Admin).Error
}

// ApproveSellerRequest implements IAdminRepository.
func (a *adminRepository) ApproveSellerRequest(userId uuid.UUID, sellerId string, approved bool) error {
	return a.db.Model(&model.Seller{}).Where("id = ?", sellerId).Updates(map[string]interface{}{
		"is_verified": approved,
		"verified_at": time.Now(),
		"verified_by": &userId,
	}).Error
}

// BlockSeller implements IAdminRepository.
func (a *adminRepository) BlockSeller(userId uuid.UUID, sellerId string, reason string) error {
	return a.db.Model(&model.Seller{}).Where("id = ?", sellerId).Updates(map[string]interface{}{
		"blocked_at":     time.Now(),
		"blocked_by":     &userId,
		"blocked_reason": reason,
	}).Error
}
func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &adminRepository{db: db}
}

func (a *adminRepository) CheckAdmin(userId uuid.UUID) (bool, error) {
	var user model.User
	err := a.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return false, err
	}
	return user.Role == enum.UserRole.Admin, nil
}

func (a *adminRepository) GetAllAdmins() ([]*model.User, error) {
	var users []*model.User
	err := a.db.Where("role = ?", enum.UserRole.Admin).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
