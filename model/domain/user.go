package domain

import "time"

//user with token
type User struct {
	Id       int
	Username string
	Password string
	Email    string
	Role     string
	Created_at time.Time
	Updated_at time.Time
	Token    string
}