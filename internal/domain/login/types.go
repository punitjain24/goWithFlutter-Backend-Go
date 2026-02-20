package login

import (
	"go-with-fiber/internal/ports/login"
)

type LoginServiceInterface interface {
	Login(user login.LoginDto) (*login.UserInfo, error)
}
