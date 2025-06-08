package notification

import (
	"ecommerce/internal/model"
	notifRepo "ecommerce/internal/repositories/notification"
	adminRepo "ecommerce/internal/repositories/user"
)

type INotificationService interface {
	NotifyAdmins(message string) error
}

type notificationService struct {
	notificationRepo notifRepo.INotificationRepository
	adminRepo        adminRepo.IAdminRepository
}

func NewNotificationService(nRepo notifRepo.INotificationRepository, aRepo adminRepo.IAdminRepository) INotificationService {
	return &notificationService{
		notificationRepo: nRepo,
		adminRepo:        aRepo,
	}
}

func (s *notificationService) NotifyAdmins(message string) error {
	admins, err := s.adminRepo.GetAllAdmins()
	if err != nil {
		return err
	}
	for _, admin := range admins {
		n := &model.Notification{
			UserID:  admin.ID,
			Message: message,
		}
		if err := s.notificationRepo.Create(n); err != nil {
			return err
		}
	}
	return nil
}
