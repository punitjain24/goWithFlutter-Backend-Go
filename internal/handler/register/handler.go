package register

import (
	registerDomain "go-with-fiber/internal/domain/register"
	"go-with-fiber/internal/ports/register"
	"go-with-fiber/internal/utility"

	"regexp"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	registerService registerDomain.RegisterInterface
	validate        *validator.Validate
}

func NewRegisterHandler(registerService registerDomain.RegisterInterface) *RegisterHandler {
	// Register custom validation
	validate := validator.New()

	//new tag registration in validator
	validate.RegisterValidation("indian_mobile", func(fl validator.FieldLevel) bool {
		mobile := fl.Field().String()

		// Indian mobile regex
		re := regexp.MustCompile(`^[6-9]\d{9}$`)

		return re.MatchString(mobile)
	})
	return &RegisterHandler{registerService: registerService, validate: validate}
}

func (r *RegisterHandler) CreateUser(ctx *fiber.Ctx) error {
	var user register.RegisterRequestDTO

	//decoding body
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utility.NewErrorResponse(
			"failed to parse the data",
			"invalid request body",
		))
	}

	// validating body

	if err := r.validate.Struct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utility.NewErrorResponse(
			"failed to validate the data",
			err.Error(),
		))
	}

	// everything is fine now we can call the service to create user
	err := r.registerService.CreateUser(&user)
	if err != nil {
		return ctx.Status(409).JSON(utility.NewErrorResponse(
			"failed to create user",
			err.Error(),
		))
	}
	//need to upgrade
	return ctx.Status(fiber.StatusCreated).JSON(utility.NewSuccessResponse(
		"user created successfully", nil,
	))
}
