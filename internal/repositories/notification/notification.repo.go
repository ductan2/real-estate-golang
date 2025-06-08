package repo

import (
	"ecommerce/internal/model"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	Create(notification *model.Notification) error
}

type notificationRepository struct {
	db *gorm.DB
}

func (r *notificationRepository) Create(notification *model.Notification) error {
	return r.db.Create(notification).Error
}

func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{db: db}
}
