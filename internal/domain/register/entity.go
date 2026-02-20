package register

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           string         `json:"id" gorm:"primaryKey"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	Country      string         `json:"country"`
	Password     string         `json:"password"`
	MobileNumber string         `json:"mobile_number" gorm:"uniqueIndex;not null"`
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
