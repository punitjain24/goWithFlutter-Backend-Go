package dbConfig

import (
	"fmt"
	"go-with-fiber/internal/config"
	"go-with-fiber/internal/domain/register"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitilizeDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect with database❌")
		return nil, err
	}
	fmt.Println("Database connected successfully ✅")
	db.AutoMigrate(&register.User{})
	return db, nil

}
