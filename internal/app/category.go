package app

import (
	"net/http"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
)

func (a *App) createCategory(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	abbreviation := r.FormValue("abbreviation")
	description := r.FormValue("description")

	_, err := a.categoryRepo.CreateCategory(userID, &domain.WorkCategory{
		Abbreviation: abbreviation,
		Description:  description,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := components.CategoryList().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
