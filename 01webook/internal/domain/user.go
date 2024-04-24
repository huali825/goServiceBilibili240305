package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string

	Nickname string
	// YYYY-MM-DD
	Birthday time.Time
	AboutMe  string

	Phone string

	Ctime time.Time
	Dtime time.Time
}
