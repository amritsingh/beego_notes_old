package models

import "time"

type User struct {
	ID        uint64 `orm:"column(id)"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
