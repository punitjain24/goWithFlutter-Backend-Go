package login

import (
	"errors"
	"go-with-fiber/internal/ports/register"
	"go-with-fiber/internal/utility"

	"gorm.io/gorm"
)

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) LoginRepository {
	return &LoginService{db: db}
}

// Login implements [LoginRepository].
func (l *LoginService) Login(login *LoginRequest) (*UserInfo, error) {
	var exisitingUser register.User

	err := l.db.Where("email = ?", login.Email).First(&exisitingUser).Error

	//check for existing user
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	//check for the password does it matches or not in hased form as we never decrypt the password
	err = utility.CheckPassword(exisitingUser.Password, login.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//created this because we dont want to expose user password
	data := &UserInfo{
		Id:        exisitingUser.Id,
		Email:     exisitingUser.Email,
		FirstName: exisitingUser.FirstName,
		LastName:  exisitingUser.LastName,
		CreatedAt: exisitingUser.CreatedAt,
		UpdatedAt: exisitingUser.UpdatedAt,
		State:     exisitingUser.State,
		Country:   exisitingUser.Country,
		City:      exisitingUser.City,
	}
	return data, nil
}
