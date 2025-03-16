package repo

import (
	"ecommerce/internal/model"
	"ecommerce/pkg/enum"
	"time"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(email string, password string, ip string, salt string) error
	VerifyOtp(email string, otp int) error
	GetUserByEmail(email string) *model.User
	UpdateUserLogin(email string, ip string) error
	GetUserById(userId string) *model.User
	UpdateUserLogout(userId string) error
	UpdateUserInfo(userId string, userInfo map[string]interface{}) error
}

type userRepository struct {
	db *gorm.DB
}

// UpdateUserInfo implements IUserRepository.
func (r *userRepository) UpdateUserInfo(userId string, userInfo map[string]interface{}) error {
	// Only update fields that are provided in the userInfo map
	return r.db.Model(&model.UserInfo{}).Where("user_id = ?", userId).Updates(userInfo).Error
}

// GetUserById implements IUserRepository.
func (r *userRepository) GetUserById(userId string) *model.User {
	var user model.User
	r.db.Preload("UserInfo").Where("id = ?", userId).First(&user)
	return &user
}

// GetUserByEmail implements IUserRepository.
func (r *userRepository) GetUserByEmail(email string) *model.User {
	var user model.User
	r.db.Where("email = ?", email).First(&user)
	return &user
}

// VerifyOtp implements IUserRepository.
func (r *userRepository) VerifyOtp(email string, otp int) error {
	return r.db.Model(&model.User{}).Where("email = ?", email).Updates(map[string]any{
		"verified":    true,
		"verified_at": time.Now(),
	}).Error
}

func (r *userRepository) CreateUser(email string, password string, ip string, salt string) error {

	return r.db.Create(&model.User{
		Email:       email,
		Password:    password,
		Role:        enum.User,
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

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}
