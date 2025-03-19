package services

import (
	"context"
	"ecommerce/global"
	"ecommerce/internal/model"
	repo "ecommerce/internal/repositories/user"

	"ecommerce/internal/utils/auth"
	"ecommerce/internal/utils/convert"
	"ecommerce/internal/utils/crypto"
	mail "ecommerce/internal/utils/email"
	"ecommerce/internal/utils/random"
	"ecommerce/internal/vo"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type IUserService interface {
	Register(username string, email string, password string, ip string, purpose *string) error
	SendOtp(email string) error
	VerifyOtp(email string, otp int) error
	Login(email string, password string, ip string) (string, error)
	GetUserInfo(userId string) *model.User
	GetUserByEmail(email string) (*model.User, error)
	Logout(userId string) error
	UpdateUserInfo(userId string, params vo.UserUpdateInfoRequest) error
}

type userService struct {
	userRepo repo.IUserRepository
}

// GetUserByEmail implements IUserService.
func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// UpdateUserInfo implements IUserService.
func (s *userService) UpdateUserInfo(userId string, params vo.UserUpdateInfoRequest) error {
	updates := convert.StructToMap(params)
	if len(updates) == 0 {
		return errors.New("nothing to update")
	}

	return s.userRepo.UpdateUserInfo(userId, updates)
}

// Logout implements IUserService.
func (s *userService) Logout(userId string) error {
	return s.userRepo.UpdateUserLogout(userId)
}

// Login implements IUserService.
func (s *userService) Login(email string, password string, ip string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if !crypto.VerifyPassword(password, user.Password, user.UserSalt) {
		return "", errors.New("password is incorrect")
	}
	err = s.userRepo.UpdateUserLogin(email, ip)
	if err != nil {
		return "", err
	}
	token, err := auth.CreateTokenJWT(&auth.PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: user.ID.String(),
		},
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyOtp implements IUserService.
func (s *userService) VerifyOtp(email string, otp int) error {
	storedOtp, err := global.Redis.Get(context.Background(), email).Result()
	if err != nil {
		return errors.New("OTP is incorrect")
	}
	if storedOtp != crypto.GetHash(strconv.Itoa(otp)) {
		return errors.New("OTP is incorrect")
	}
	err = s.userRepo.VerifyOtp(email, otp)
	if err != nil {
		return err
	}
	// delete otp from redis when verify success
	err = global.Redis.Del(context.Background(), email).Err()
	if err != nil {
		return err
	}
	return nil
}

// SendOtp implements IUserService.
func (s *userService) SendOtp(email string) error {
	// send otp to email and save otp to db and expire time 10 minutes
	otp := random.GenerateSixDigitOtp()
	hashedOtp := crypto.GetHash(strconv.Itoa(otp))
	err := global.Redis.Set(context.Background(), email, hashedOtp, 10*time.Minute).Err()
	if err != nil {
		// logger
		return err
	}

	err = mail.SendEmailOtp([]string{email}, global.Config.SMTP.Username, otp)
	if err != nil {
		return err
	}
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	return nil
}

func (s *userService) Register(username string, email string, password string, ip string, purpose *string) error {
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	return s.userRepo.CreateUser(username, email, hashedPassword, ip, salt)
}
func (s *userService) GetUserInfo(userId string) *model.User {
	return s.userRepo.GetUserById(userId)
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}
