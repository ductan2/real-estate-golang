package initialize

import (
	"ecommerce/global"
	"ecommerce/internal/model"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	// Default database configuration
	db := global.Config.Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.Username, db.Password, db.Dbname)


	// Create GORM config with logging
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Open database with GORM and PostgreSQL
	gormDB, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		global.Logger.Errorf("failed to open database: %w", err)
		return nil, err
	}

	// Get the underlying *sql.DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		global.Logger.Errorf("failed to get database instance: %w", err)
		return nil, err
	}

	// Set default connection pool settings
	sqlDB.SetMaxOpenConns(db.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(db.MaxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(db.MaxLifetimeConnection) * time.Second)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		global.Logger.Errorf("failed to ping database: %w", err)
		return nil, err
	}

	// Auto migrate models
	if err := AutoMigrate(gormDB); err != nil {
		global.Logger.Errorf("failed to migrate models: %w", err)
		return nil, err
	}
	global.Logger.Infof("Database migrated successfully")

	global.DB = gormDB
	return gormDB, nil
}

func AutoMigrate(db *gorm.DB) error {
	// Init tables here
	return db.AutoMigrate(model.AllModels()...)
}
