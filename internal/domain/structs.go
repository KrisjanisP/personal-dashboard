package domain

import "time"

type User struct {
	ID        int32
	Username  string
	Password  string
	CreatedAt time.Time
}

type WorkCategory struct {
	ID           int32
	OwnerUserID  int32
	Abbreviation string
	Description  string
}

type TimeEntry struct {
	ID            int32
	OwnerUserID   int32
	CategoryID    int32
	StartDateTime time.Time
	EndDateTime   time.Time
}
