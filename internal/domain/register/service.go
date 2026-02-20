package register

import (
	"errors"
	"go-with-fiber/internal/ports/register"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	registerRepo register.RegisterRepository
}

func NewRegisterService(register register.RegisterRepository) RegisterInterface {
	return &RegisterService{registerRepo: register}
}

// CreateUser implements [RegisterInterface].
func (r *RegisterService) CreateUser(userDto *register.RegisterRequestDTO) error {

	//validation part
	if userDto.Password != userDto.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	//if everything matched than encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// password encryption is done now convert to the user model

	user := &register.User{
		FirstName:    userDto.FirstName,
		LastName:     userDto.LastName,
		Email:        userDto.Email,
		City:         userDto.City,
		State:        userDto.State,
		Country:      userDto.Country,
		Password:     string(hashedPassword),
		MobileNumber: userDto.MobileNumber,
	}

	//save to Db
	return r.registerRepo.CreateUser(user)
}
