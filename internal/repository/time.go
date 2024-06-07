package repository

import (
	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/database/model"
	"github.com/KrisjanisP/personal-dashboard/internal/database/table"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/go-jet/jet/v2/sqlite"
	"github.com/jmoiron/sqlx"
)

type timeEntryRepoImpl struct {
	db *sqlx.DB
}

// CreateTimeEntry implements internal.TimeEntryRepo.
func (t *timeEntryRepoImpl) CreateTimeEntry(timeEntry *domain.TimeEntry) (int32, error) {
	stmt := table.TimeTracking.INSERT(table.TimeTracking.MutableColumns).
		MODEL(&model.TimeTracking{
			ID:         new(int32),
			UserID:     timeEntry.OwnerUserID,
			CategoryID: timeEntry.CategoryID,
			StartTime:  timeEntry.StartDateTime,
			EndTime:    &timeEntry.EndDateTime,
		}).RETURNING(table.TimeTracking.ID)

	var record model.TimeTracking
	err := stmt.Query(t.db, &record)
	if err != nil {
		return 0, err
	}

	return *record.ID, nil
}

// DeleteTimeEntry implements internal.TimeEntryRepo.
func (t *timeEntryRepoImpl) DeleteTimeEntry(timeEntryID int32) error {
	stmt := table.TimeTracking.DELETE().
		WHERE(table.TimeTracking.ID.EQ(sqlite.Int32(timeEntryID)))

	_, err := stmt.Exec(t.db)
	if err != nil {
		return err
	}
	return nil
}

// GetTimeEntryByID implements internal.TimeEntryRepo.
func (t *timeEntryRepoImpl) GetTimeEntryByID(timeEntryID int32) (*domain.TimeEntry, error) {
	stmt := sqlite.SELECT(table.TimeTracking.AllColumns).
		FROM(table.TimeTracking).
		WHERE(table.TimeTracking.ID.EQ(sqlite.Int32(timeEntryID)))

	var record model.TimeTracking
	err := stmt.Query(t.db, &record)
	if err != nil {
		return nil, err
	}

	return &domain.TimeEntry{
		ID:            *record.ID,
		OwnerUserID:   record.UserID,
		CategoryID:    record.CategoryID,
		StartDateTime: record.StartTime,
		EndDateTime:   *record.EndTime,
	}, nil
}

// ListTimeEntries implements internal.TimeEntryRepo.
func (t *timeEntryRepoImpl) ListTimeEntries(userID int32) ([]*domain.TimeEntry, error) {
	stmt := sqlite.SELECT(table.TimeTracking.AllColumns).
		FROM(table.TimeTracking).
		WHERE(table.TimeTracking.UserID.EQ(sqlite.Int32(userID)))

	var records []model.TimeTracking
	err := stmt.Query(t.db, &records)
	if err != nil {
		return nil, err
	}

	var mapped []*domain.TimeEntry
	for _, record := range records {
		mapped = append(mapped, &domain.TimeEntry{
			ID:            *record.ID,
			OwnerUserID:   record.UserID,
			CategoryID:    record.CategoryID,
			StartDateTime: record.StartTime,
			EndDateTime:   *record.EndTime,
		})
	}

	return mapped, nil
}

func NewTimeEntryRepository(db *sqlx.DB) *timeEntryRepoImpl {
	return &timeEntryRepoImpl{db: db}
}

var _ internal.TimeEntryRepo = &timeEntryRepoImpl{}
