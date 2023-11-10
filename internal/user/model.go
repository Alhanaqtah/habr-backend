package user

import "time"

type User struct {
	Id               int       `json:"id"`
	Username         string    `json:"username"`
	RegistrationDate time.Time `json:"registration_date"`
	Rating           int       `json:"rating"`
	Karma            int       `json:"karma"`
	Publications     []int     `json:"publications"`
	Bookmarks        []int     `json:"bookmarks"`
	Followers        []int     `json:"followers"`
	Followings       []int     `json:"followings"`
}
