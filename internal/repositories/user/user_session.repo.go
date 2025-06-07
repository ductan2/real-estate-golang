package repo

import (
	"ecommerce/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserSessionRepository interface {
	CreateUserSession(userId uuid.UUID, ipAddress string, location string, device string, userAgent string) error
	GetUserSessionByUserId(userId uuid.UUID) (*model.UserSession, error)
	DeleteUserSession(userId uuid.UUID) error
}

type userSessionRepository struct {
	db *gorm.DB
}

func NewUserSessionRepository(db *gorm.DB) IUserSessionRepository {
	return &userSessionRepository{db: db}
}

func (u *userSessionRepository) CreateUserSession(userId uuid.UUID, ipAddress string, location string, device string, userAgent string) error {
	return u.db.Create(&model.UserSession{
		UserId:    userId,
		IpAddress: ipAddress,
		Location:  location,
		Device:    device,
		UserAgent: userAgent,
	}).Error
}

func (u *userSessionRepository) GetUserSessionByUserId(userId uuid.UUID) (*model.UserSession, error) {
	var userSession model.UserSession
	err := u.db.Where("user_id = ?", userId).First(&userSession).Error
	if err != nil {
		return nil, err
	}
	return &userSession, nil
}

func (u *userSessionRepository) DeleteUserSession(userId uuid.UUID) error {
	return u.db.Where("user_id = ?", userId).Delete(&model.UserSession{}).Error
}