package internal

import "github.com/KrisjanisP/personal-dashboard/internal/domain"

type UserRepo interface {
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(user *domain.User) (int32, error)
	GetUserByID(userID int32) (*domain.User, error)
}

type CategoryRepo interface {
	GetCategoryByID(userID int32, categoryID int32) (*domain.WorkCategory, error)
	CreateCategory(userID int32, category *domain.WorkCategory) (int32, error)
	ListCategories(userID int32) ([]*domain.WorkCategory, error)
	DeleteCategory(userID int32, categoryID int32) error
}
