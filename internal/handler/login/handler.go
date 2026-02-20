package login

import (
	lDomain "go-with-fiber/internal/domain/login"
	lPorts "go-with-fiber/internal/ports/login"
	"go-with-fiber/internal/utility"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
	loginService lDomain.LoginServiceInterface
	validate     *validator.Validate
}

func NewLoginHandler(loginService lDomain.LoginServiceInterface) *LoginHandler {
	return &LoginHandler{loginService: loginService, validate: validator.New()}
}

func (l *LoginHandler) Login(ctx *fiber.Ctx) error {
	var user lPorts.LoginDto

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utility.NewErrorResponse(
			"failed to parse the data",
			"invalid request body",
		))
	}

	//validating struct
	if err := l.validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utility.NewErrorResponse(
			"validation failed",
			err.Error(),
		))
	}

	//calling service layer 
	data, err := l.loginService.Login(user)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(utility.NewErrorResponse(
			"failed to login",
			err.Error(),
		))
	}

	//creating jwt token for the authorization
	token, err := utility.GenerateJWT(data.Id, data.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utility.NewErrorResponse(
			"failed to generate token",
			err.Error(),
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(utility.NewSuccessResponse(
		"login successfully",
		fiber.Map{
			"user_info": data,
			"token":     token,
		},
	))

}
