package user

type storage interface {
	GetUsers() *[]User
	GetUser(username string) *User
	GetUserPublications(username string) *User
	GetUserFollowers(username string) *User
	GetUserFollowings(username string) *User
}
