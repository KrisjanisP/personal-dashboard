package app

import (
	"net/http"

	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	Addr string
}

func NewApp(addr string) *App {
	return &App{Addr: addr}
}

func (a *App) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", a.Home)
	http.ListenAndServe(a.Addr, r)
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	if 2%2 == 0 {
		if err := pages.AuthenticationPage().Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	user := domain.User{
		ID:       0,
		Username: "hello",
	}
	if err := pages.HomePage(user).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
