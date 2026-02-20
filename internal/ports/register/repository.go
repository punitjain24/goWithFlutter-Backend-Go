package register

type RegisterRepository interface {
	CreateUser(user *User) error
}
