package app

import (
	"context"
	"errors"
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
	r.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("web/public"))).ServeHTTP)

	r.Group(func(r chi.Router) {
		r.Use(a.sessionManager.LoadAndSave)
		r.Use(a.Slow)
		r.Get("/login", a.LoginGet)
		r.Put("/login", a.LoginPut)
		r.Put("/logout", a.LogoutPut)
		r.With(a.AuthMiddleware).Get("/", a.Home)
	})
	http.ListenAndServe(a.Addr, r)
}

func (a *App) Slow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		next.ServeHTTP(w, r)
	})
}

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

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)
	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		if err := pages.ErrorPage("internal server error").Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := pages.HomePage(user).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) LoginGet(w http.ResponseWriter, r *http.Request) {
	if err := pages.AuthenticationPage(nil).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) LoginPut(w http.ResponseWriter, r *http.Request) {
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
	if err := pages.HomePage(user).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *App) LogoutPut(w http.ResponseWriter, r *http.Request) {
	a.sessionManager.Remove(r.Context(), "user_id")
	w.Header().Set("HX-Redirect", "/login")
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}
