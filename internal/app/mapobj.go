package app

import (
	"fmt"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
)

func (a *App) mapDomainTimeEntryToTimeTrackTableRow(t *domain.TimeEntry) (*components.TimeTrackerHistoryTableRow, error) {
	res := &components.TimeTrackerHistoryTableRow{
		TimeEntryID:  int(t.ID),
		CategoryAbbr: "",
		StartTime:    "",
		EndTime:      "",
		Duration:     "",
	}

	category, err := a.categoryRepo.GetCategoryByID(t.CategoryID)
	if err != nil {
		return nil, err
	}

	res.CategoryAbbr = category.Abbreviation
	startTimeMarshalled := t.StartDateTime.Local().Format("2006-01-02 15:04:05")

	res.StartTime = string(startTimeMarshalled)

	endTimeMarshalled := t.EndDateTime.Local().Format("2006-01-02 15:04:05")

	res.EndTime = string(endTimeMarshalled)

	seconds := t.EndDateTime.Sub(t.StartDateTime).Seconds()
	res.Duration = fmt.Sprintf("%02d:%02d:%02d", int(seconds/3600), int(seconds/60)%60, int(seconds)%60)

	return res, nil
}
