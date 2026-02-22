package dashboard

import (
	"fmt"
	"go-with-fiber/internal/config"
	"go-with-fiber/internal/domain/dashboard"
	"go-with-fiber/internal/utility"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	dashboardService dashboard.DashboardServiceInterface
	cfg              config.Config
}

func NewDashboardHandler(dashboardService dashboard.DashboardServiceInterface, cfg config.Config) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService, cfg: cfg}
}

func (d *DashboardHandler) FetchUserList(ctx *fiber.Ctx) error {
	userId, ok := ctx.Locals("user_id").(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			utility.NewErrorResponse("Unauthorized", "Invalid token"),
		)
	}

	//this is for testing purpose will remove this in production
	fmt.Println("my test user id,", userId)
	limit := ctx.QueryInt("limit", 10)
	pageIndex := ctx.QueryInt("pageIndex", 1)

	// final calling service
	userList, err := d.dashboardService.FetchUserList(limit, pageIndex)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(utility.NewErrorResponse("Unable to fetch the data", err.Error()))
	}

	//success response
	return ctx.Status(fiber.StatusOK).JSON(
		utility.NewSuccessResponse(
			"Data fetched successfully",
			userList,
		),
	)

}
