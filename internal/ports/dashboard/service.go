package dashboard

import (
	"go-with-fiber/internal/ports/register"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) DashboardRepostiory {
	return &DashboardService{db: db}
}

// FetchUserList implements [DashboardRepostiory] with pagination
func (d *DashboardService) FetchUserList(limit, pageIndex int) ([]UserList, error) {
	var data []UserList

	// we need to fetch data in pagnated format which will not impact our services for large data
	if limit == 0 {
		limit = 10
	}

	if pageIndex <= 0 {
		pageIndex = 1
	}
	offset := (pageIndex - 1) * limit
	err := d.db.Model(&register.User{}).Limit(limit).Offset(offset).
		Select("first_name", "last_name", "email").Order("created_at desc").
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
