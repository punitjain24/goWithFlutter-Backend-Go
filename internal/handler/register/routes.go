package register

import "github.com/gofiber/fiber/v2"

func (r *RegisterHandler) RegisterRoute(router fiber.Router) {
	router.Post("/register", r.CreateUser)
}
