package app

import (
	"net/http"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
)

func (a *App) createCategory(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoryName := r.FormValue("name")
	userID := r.Context().Value("user_id").(int32)

	_, err := a.categoryRepo.CreateCategory(&domain.WorkCategory{
		ID:   userID,
		Name: categoryName,
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
