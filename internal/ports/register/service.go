package register

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type RegisterService struct {
	db *gorm.DB
}

func NewRegisterService(db *gorm.DB) RegisterRepository {
	return &RegisterService{db: db}
}

// CreateUser implements [RegisterRepository].
func (s *RegisterService) CreateUser(user *User) error {
	err := s.db.Create(user).Error
	if err != nil {

		// Check duplicate key error
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("email or mobile number already exists")
		}

		// Return actual DB error
		return err
	}

	return nil
}
