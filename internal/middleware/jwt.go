package middleware

import (
	"fmt"
	"go-with-fiber/internal/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(jwtSecret string) fiber.Handler {

	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Token not found",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid token format",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("mty jwt secret", jwtSecret)
		userID, err := utility.ParseJWT(tokenString, jwtSecret)
		fmt.Print(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		// Store userID in context
		c.Locals("user_id", userID)

		return c.Next()
	}
}
