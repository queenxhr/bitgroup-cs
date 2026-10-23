package database

import (
	"fmt"
	"log"

	"bitgroup-cs/config"
	categoryRepo "bitgroup-cs/infrastructure/repository/category"
	productRepo "bitgroup-cs/infrastructure/repository/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgresql database: %w", err)
	}

	log.Println("Database connection established successfully")

	// Run auto migrations
	if err := AutoMigrate(db); err != nil {
		log.Printf("Warning: AutoMigrate encountered an issue: %v\n", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(
		&categoryRepo.CategoryModel{},
		&productRepo.ProductModel{},
	)
	if err != nil {
		return fmt.Errorf("automigration failed: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}
