package dashboard

type DashboardRepostiory interface {
	FetchUserList(limit, pageIndex int) ([]UserList, error)
}
