package repository

import (
	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
)

type categoryRepoImpl struct {
}

// CreateCategory implements internal.CategoryRepo.
func (c *categoryRepoImpl) CreateCategory(category *domain.WorkCategory) (int32, error) {
	// TODO: implement
	panic("unimplemented")
}

var _ internal.CategoryRepo = &categoryRepoImpl{}
