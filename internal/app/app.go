package app

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/domain"
	"github.com/KrisjanisP/personal-dashboard/internal/repository"
	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	Addr           string
	sessionManager *scs.SessionManager
	userRepo       internal.UserRepo
}

func NewApp(addr string) *App {
	app := &App{Addr: addr}

	app.sessionManager = scs.New()
	app.sessionManager.Lifetime = 24 * time.Hour

	sqliteDB := sqlx.MustConnect("sqlite3", "./data/sqlite3.db")
	app.userRepo = repository.NewUserRepository(sqliteDB)

	return app
}

func (a *App) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", a.Home)
	r.Put("/login", a.Login)
	http.ListenAndServe(a.Addr, r)
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	// msg := sessionManager.GetString(r.Context(), "message")
	if 2%2 == 0 { // this will be replaces by auth middleware
		var errMsg string = fmt.Sprintf("Error: %s.", "some error here")
		if err := pages.AuthenticationPage(&errMsg).Render(r.Context(), w); err != nil {
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

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

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

	a.sessionManager.Put(r.Context(), "user_id", user.ID)

	var errMsg string = "further actions are not implemented"
	if err := pages.AuthenticationPage(&errMsg).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) Logout(w http.ResponseWriter, r *http.Request) {
	// sessionManager.Remove(r.Context(), "message")
}
