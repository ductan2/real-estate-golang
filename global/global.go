package global

import (
	"ecommerce/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	DB *gorm.DB
	Redis *redis.Client
)
