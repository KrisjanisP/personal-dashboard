package repository

import (
	"errors"

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

// GetCategoryByID implements internal.CategoryRepo.
func (c *categoryRepoImpl) GetCategoryByID(userID int32, categoryID int32) (*domain.WorkCategory, error) {
	stmt := sqlite.SELECT(table.WorkCategories.AllColumns).
		FROM(table.WorkCategories).
		WHERE(table.WorkCategories.UserID.EQ(sqlite.Int32(userID)).AND(table.WorkCategories.ID.EQ(sqlite.Int32(categoryID))))

	var record model.WorkCategories
	err := stmt.Query(c.db, &record)
	if err != nil {
		return nil, err
	}

	return &domain.WorkCategory{
		ID:           *record.ID,
		OwnerUserID:  *record.UserID,
		Abbreviation: *record.Abbreviation,
		Description:  *record.Description,
	}, nil
}

// DeleteCategory implements internal.CategoryRepo.
func (c *categoryRepoImpl) DeleteCategory(userID int32, categoryID int32) error {
	// mark deleted column as 1
	stmt := table.WorkCategories.UPDATE(table.WorkCategories.Deleted).SET(sqlite.Int32(1)).
		WHERE(table.WorkCategories.UserID.EQ(sqlite.Int32(userID)).AND(table.WorkCategories.ID.EQ(sqlite.Int32(categoryID))))

	_, err := stmt.Exec(c.db)
	if err != nil {
		return err
	}
	return nil
}

// ListCategories implements internal.CategoryRepo.
func (c *categoryRepoImpl) ListCategories(userID int32) ([]*domain.WorkCategory, error) {
	stmt := sqlite.SELECT(table.WorkCategories.AllColumns).
		FROM(table.WorkCategories).
		WHERE(table.WorkCategories.UserID.EQ(sqlite.Int32(userID)).AND(table.WorkCategories.Deleted.EQ(sqlite.Int32(0))))

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
	if category == nil {
		return 0, errors.New("category is nil")
	}
	if category.Abbreviation == "" {
		return 0, errors.New("abbreviation is empty")
	}
	if category.Description == "" {
		return 0, errors.New("description is empty")
	}
	deleted := int32(0)
	stmt := table.WorkCategories.INSERT(table.WorkCategories.MutableColumns).
		MODEL(&model.WorkCategories{
			UserID:       &userID,
			Abbreviation: &category.Abbreviation,
			Description:  &category.Description,
			Deleted:      &deleted,
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
