package app

import (
	"net/http"
	"strconv"

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
	
	if err := components.StopTimeComponent(category).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
