package dashboard

import "go-with-fiber/internal/ports/dashboard"

type DashboardService struct {
	dashboardRepo dashboard.DashboardRepostiory
}

func NewDashboardService(dashboardRepo dashboard.DashboardRepostiory) DashboardServiceInterface {
	return &DashboardService{dashboardRepo: dashboardRepo}
}

// FetchUserList implements [DashboardServiceInterface].
func (d *DashboardService) FetchUserList(limit int, pageIndex int) ([]dashboard.UserList, error) {
	return d.dashboardRepo.FetchUserList(limit, pageIndex)
}
