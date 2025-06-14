package repo

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/enum"
	"errors"
	"time"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(username string, email string, password string, ip string, salt string) error
	VerifyOtp(email string, otp int) error
	GetUserByEmail(email string) (*model.User, error)
	UpdateUserLogin(email string, ip string) error
	GetUserById(userId string) *model.User
	UpdateUserLogout(userId string) error
	UpdateUserInfo(userId string, userInfo map[string]interface{}) error
	UpdateUserAvatar(userId string, avatarUrl string) error
}

type userRepository struct {
	db *gorm.DB
}

// UpdateUserInfo implements IUserRepository.
func (r *userRepository) UpdateUserInfo(userId string, userInfo map[string]interface{}) error {
	// Only update fields that are provided in the userInfo map
	if userInfo["username"] != nil {
		r.db.Model(&model.User{}).Where("id = ?", userId).Update("username", userInfo["username"])
	}
	// remove username from userInfo
	delete(userInfo, "username")
	return r.db.Model(&model.UserInfo{}).Where("user_id = ?", userId).Updates(userInfo).Error
}

// GetUserById implements IUserRepository.
func (r *userRepository) GetUserById(userId string) *model.User {
	var user model.User
	r.db.Preload("UserInfo").Where("id = ?", userId).First(&user)
	return &user
}

// GetUserByEmail implements IUserRepository.
func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// VerifyOtp implements IUserRepository.
func (r *userRepository) VerifyOtp(email string, otp int) error {
	return r.db.Model(&model.User{}).Where("email = ?", email).Updates(map[string]any{
		"verified":    true,
		"verified_at": time.Now(),
	}).Error
}

func (r *userRepository) CreateUser(username string, email string, password string, ip string, salt string) error {

	return r.db.Create(&model.User{
		Username:    username,
		Email:       email,
		Password:    password,
		Role:        enum.UserRole.User,
		UserSalt:    salt,
		UserLoginIP: ip,
	}).Error
}

func (r *userRepository) UpdateUserLogin(email string, ip string) error {
	return r.db.Model(&model.User{}).Where("email = ?", email).Updates(map[string]any{
		"user_login_ip":   ip,
		"user_login_time": time.Now(),
	}).Error
}

func (r *userRepository) UpdateUserLogout(userId string) error {
	return r.db.Model(&model.User{}).Where("id = ?", userId).
		Update("user_logout_time", time.Now()).
		Error
}

func (s *userRepository) UpdateUserAvatar(userId string, avatarUrl string) error {
	return s.db.Model(&model.UserInfo{}).Where("user_id = ?", userId).Update("avatar", avatarUrl).Error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}
