package global

import (
	"ecommerce/internal/services/queue"
	"ecommerce/pkg/setting"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	RabbitMQ *queue.QueueService
	DB     *gorm.DB
	Redis  *redis.Client
	Logger *logrus.Logger
)
