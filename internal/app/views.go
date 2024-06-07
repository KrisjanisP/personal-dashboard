package app

import (
	"net/http"

	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
)

func (a *App) renderHomeView(w http.ResponseWriter, r *http.Request) {
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

	if err := pages.HomePage(user, categories, timeEntries).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
