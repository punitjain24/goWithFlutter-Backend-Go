package login

import (
	"go-with-fiber/internal/ports/login"
)

type LoginService struct {
	loginRepo login.LoginRepository
}

func NewLoginService(loginRepo login.LoginRepository) LoginServiceInterface {
	return &LoginService{loginRepo: loginRepo}
}

// Login implements [LoginServiceInterface].
func (l *LoginService) Login(user login.LoginDto) (*login.UserInfo, error) {
	loginReq := &login.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	return l.loginRepo.Login(loginReq)
}
