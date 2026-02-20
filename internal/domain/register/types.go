package register

import "go-with-fiber/internal/ports/register"

type RegisterInterface interface {
	CreateUser(user *register.RegisterRequestDTO) error
}
