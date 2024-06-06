package internal

import "github.com/KrisjanisP/personal-dashboard/internal/domain"

type UserRepo interface {
	GetUserByUsername(username string) (*domain.User, error)
}
