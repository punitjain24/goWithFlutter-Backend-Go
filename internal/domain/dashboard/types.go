package dashboard

import (
	"go-with-fiber/internal/ports/dashboard"
)

type DashboardServiceInterface interface {
	FetchUserList(limit, pageIndex int) ([]dashboard.UserList, error)
}
