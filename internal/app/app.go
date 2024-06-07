package app

import (
	"log"
	"net/http"
	"time"

	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/repository"
	"github.com/alexedwards/scs/sqlite3store"
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
	categoryRepo   internal.CategoryRepo
	timeEntryRepo  internal.TimeEntryRepo
}

func NewApp(addr string) *App {
	app := &App{Addr: addr}

	sqliteDB := sqlx.MustConnect("sqlite3", "./data/sqlite3.db")

	app.sessionManager = scs.New()
	app.sessionManager.Store = sqlite3store.New(sqliteDB.DB)
	app.sessionManager.Lifetime = 24 * time.Hour

	app.userRepo = repository.NewUserRepository(sqliteDB)
	app.categoryRepo = repository.NewCategoryRepository(sqliteDB)
	app.timeEntryRepo = repository.NewTimeEntryRepository(sqliteDB)

	return app
}

func (a *App) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("web/public"))).ServeHTTP)

	// routes that require auth cookie
	r.Group(func(r chi.Router) {
		r.Use(a.sessionManager.LoadAndSave)
		// r.Use(a.Slow)
		r.Get("/login", a.viewAuthPage)

		r.Put("/login", a.login)
		r.Put("/logout", a.logout)

		r.Group(func(r chi.Router) {
			r.Use(a.AuthMiddleware)
			r.Put("/category", a.createCategory)
			r.Delete("/category/{id}", a.deleteCategory)
			r.Get("/time/start", a.startTime)
			r.Put("/time/stop", a.stopTime)
			r.Get("/", a.renderHome)
		})
	})
	log.Println("Listening on", a.Addr)
	http.ListenAndServe(a.Addr, r)
}

func (a *App) Slow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		next.ServeHTTP(w, r)
	})
}
