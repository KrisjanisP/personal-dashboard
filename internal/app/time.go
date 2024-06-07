package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
)

func (a *App) startTime(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoryIDStr := r.FormValue("work-category")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := a.categoryRepo.GetCategoryByID(int32(categoryID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()
	marshalled, err := currentTime.MarshalText()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := components.StopTimeComponent(category, string(marshalled)).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) stopTime(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startStr := r.FormValue("start")
	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeInSecondsStr := r.FormValue("seconds")
	timeInSeconds, err := strconv.Atoi(timeInSecondsStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryStr := r.FormValue("work-category")
	category, err := a.categoryRepo.GetCategoryByAbbreviation(userID, categoryStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	_, err = a.timeEntryRepo.CreateTimeEntry(&domain.TimeEntry{
		OwnerUserID:   userID,
		CategoryID:    category.ID,
		StartDateTime: start,
		EndDateTime:  	start.Add(time.Duration(timeInSeconds) * time.Second),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.renderHome(w, r)
}
