package login

import (
	"github.com/gofiber/fiber/v2"
)

func (l *LoginHandler) LoginRoute(router fiber.Router) {
	router.Post("/login", l.Login)

}
