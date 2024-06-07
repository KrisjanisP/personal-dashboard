package app

import (
	"log"
	"net/http"
	"sort"

	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
)

func (a *App) renderHome(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)

	categories, err := a.categoryRepo.ListCategories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	timeEntries, err := a.timeEntryRepo.ListTimeEntries(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// sort by start time (oldest comes first)
	sort.Slice(timeEntries, func(i, j int) bool {
		return timeEntries[i].StartDateTime.After(timeEntries[j].StartDateTime)
	})

	timeTrackerTableRows := make([]*components.TimeTrackerHistoryTableRow, 0)
	for _, t := range timeEntries {
		row, err := a.mapDomainTimeEntryToTimeTrackTableRow(t)
		if err != nil {
			log.Println("error mapping domain time entry to time tracker table row:", err)
			if err := pages.ErrorPage("internal server error").Render(r.Context(), w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		timeTrackerTableRows = append(timeTrackerTableRows, row)
	}

	if err := pages.HomePage(user, categories, timeTrackerTableRows).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
