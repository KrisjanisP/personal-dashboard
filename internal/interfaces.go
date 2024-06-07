package internal

import "github.com/KrisjanisP/personal-dashboard/internal/domain"

type UserRepo interface {
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(user *domain.User) (int32, error)
	GetUserByID(userID int32) (*domain.User, error)
}

type CategoryRepo interface {
	GetCategoryByID(categoryID int32) (*domain.WorkCategory, error)
	GetCategoryByAbbreviation(userID int32, abbreviation string) (*domain.WorkCategory, error)
	CreateCategory(category *domain.WorkCategory) (int32, error)
	ListCategories(userID int32) ([]*domain.WorkCategory, error)
	DeleteCategory(userID int32, categoryID int32) error
}

type TimeEntryRepo interface {
	GetTimeEntryByID(timeEntryID int32) (*domain.TimeEntry, error)
	CreateTimeEntry(timeEntry *domain.TimeEntry) (int32, error)
	ListTimeEntries(userID int32) ([]*domain.TimeEntry, error)
	DeleteTimeEntry(timeEntryID int32) error
}
