package repository

import (
	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/database/model"
	"github.com/KrisjanisP/personal-dashboard/internal/database/table"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/go-jet/jet/v2/sqlite"
	"github.com/jmoiron/sqlx"
)

type categoryRepoImpl struct {
	db *sqlx.DB
}

// ListCategories implements internal.CategoryRepo.
func (c *categoryRepoImpl) ListCategories(userID int32) ([]*domain.WorkCategory, error) {
	stmt := sqlite.SELECT(table.WorkCategories.AllColumns).
		FROM(table.WorkCategories).
		WHERE(table.WorkCategories.UserID.EQ(sqlite.Int32(userID)))

	var records []model.WorkCategories
	err := stmt.Query(c.db, &records)
	if err != nil {
		return nil, err
	}

	var mapped []*domain.WorkCategory
	for _, record := range records {
		mapped = append(mapped, &domain.WorkCategory{
			ID:           *record.ID,
			Abbreviation: *record.Abbreviation,
			Description:  *record.Description,
		})
	}

	return mapped, nil
}

// CreateCategory implements internal.CategoryRepo.
func (c *categoryRepoImpl) CreateCategory(userID int32, category *domain.WorkCategory) (int32, error) {
	stmt := table.WorkCategories.INSERT(table.WorkCategories.MutableColumns).
		MODEL(&model.WorkCategories{
			UserID:       &userID,
			Abbreviation: &category.Abbreviation,
			Description:  &category.Description,
		}).RETURNING(table.WorkCategories.ID)

	var record model.WorkCategories
	err := stmt.Query(c.db, &record)
	if err != nil {
		return 0, err
	}

	return *record.ID, nil
}

func NewCategoryRepository(db *sqlx.DB) *categoryRepoImpl {
	return &categoryRepoImpl{
		db: db,
	}
}

var _ internal.CategoryRepo = &categoryRepoImpl{}
