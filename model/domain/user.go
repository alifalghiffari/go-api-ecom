package domain

import "time"

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	Role       bool
	Created_at time.Time
	Updated_at time.Time
	Token      string
}
