package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/model"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Default database configuration
	db := global.Config.Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.Username, db.Password, db.Dbname)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// Set default connection pool settings
	sqlDB.SetMaxOpenConns(db.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(db.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(db.MaxLifetimeConnection) * time.Second)

	// Auto migrate models
	if err := AutoMigrate(gormDB); err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}
	fmt.Println("Database migrated successfully")
	global.DB = gormDB
	return gormDB, nil
}

func AutoMigrate(db *gorm.DB) error {
	// Init tables here
	return db.AutoMigrate(model.AllModels()...)
}

