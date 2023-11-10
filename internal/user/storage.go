package user

type storage interface {
	GetUsers() *[]User
}
