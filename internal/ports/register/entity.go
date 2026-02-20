package register

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           string         `json:"id" gorm:"primaryKey"`
	FirstName    string         `json:"first_name" validate:"required"`
	LastName     string         `json:"last_name" validate:"required"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	City         string         `json:"city" validate:"required"`
	State        string         `json:"state" validate:"required"`
	Country      string         `json:"country" validate:"required"`
	Password     string         `json:"password" validate:"required,min=6"`
	MobileNumber string         `json:"mobile_number" gorm:"uniqueIndex;not null" validate:"required,indian_mobile"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Id == "" {
		u.Id = "user_" + uuid.New().String()
	}
	return
}
