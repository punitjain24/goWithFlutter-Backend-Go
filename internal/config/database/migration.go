package dbConfig

import (
	"go-with-fiber/internal/ports/register"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(
		&register.User{},
	)
}
