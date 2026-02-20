package login

type LoginRepository interface {
	Login(login *LoginRequest) (*UserInfo, error)
}
