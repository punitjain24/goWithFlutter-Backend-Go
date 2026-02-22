package dashboard

import (
	"go-with-fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (d *DashboardHandler) DashboardRoutes(router fiber.Router) {

	//path creation with middleware
	protected := router.Group("/", middleware.JWTMiddleware(d.cfg.JWTSecret))

	//api path
	protected.Get("/dashboard", d.FetchUserList)
}
