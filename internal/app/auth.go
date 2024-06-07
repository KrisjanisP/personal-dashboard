package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
)

func (a *App) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := a.sessionManager.GetInt32(r.Context(), "user_id")
		if userID == 0 {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *App) viewAuthPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.AuthenticationPage(nil).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			var errMsg string = "invalid username or password"
			if err := pages.AuthenticationPage(&errMsg).Render(r.Context(), w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		var errMsg string = "internal server error"
		if err := pages.AuthenticationPage(&errMsg).Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if user.Password != password {
		var errMsg string = "invalid username or password"
		if err := pages.AuthenticationPage(&errMsg).Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	a.sessionManager.Put(r.Context(), "user_id", user.ID)

	w.Header().Set("HX-Push-Url", "/")
	ctx := context.WithValue(r.Context(), "user_id", user.ID)
	r = r.WithContext(ctx)
	a.renderHome(w, r)
}

func (a *App) logout(w http.ResponseWriter, r *http.Request) {
	a.sessionManager.Remove(r.Context(), "user_id")
	w.Header().Set("HX-Push-Url", "/login")
	if err := pages.AuthenticationPage(nil).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
