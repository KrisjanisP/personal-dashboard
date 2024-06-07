package repository

import (
	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/jmoiron/sqlx"
)

type categoryRepoImpl struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) internal.CategoryRepo {
	return &categoryRepoImpl{
		db: db,
	}
}

// CreateCategory implements internal.CategoryRepo.
func (c *categoryRepoImpl) CreateCategory(category *domain.WorkCategory) (int32, error) {
	// TODO: implement
	panic("unimplemented")
}

var _ internal.CategoryRepo = &categoryRepoImpl{}
