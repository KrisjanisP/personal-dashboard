package domain

import "time"

type User struct {
	ID        int32
	Username  string
	Password  string
	CreatedAt time.Time
}
