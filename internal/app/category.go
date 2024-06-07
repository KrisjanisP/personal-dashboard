package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/components"
	"github.com/go-chi/chi"
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

	categories, err := a.categoryRepo.ListCategories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := components.CategoryList(categories).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) deleteCategory(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)

	categoryIDStr := chi.URLParam(r, "id")
	log.Println("categoryIDStr", categoryIDStr)
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := a.categoryRepo.GetCategoryByID(userID, int32(categoryID))
	if err != nil || category == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if category.OwnerUserID != userID {
		http.Error(w, "category does not belong to user", http.StatusForbidden)
		return
	}

	err = a.categoryRepo.DeleteCategory(userID, int32(categoryID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := a.categoryRepo.ListCategories(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := components.CategoryList(categories).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
