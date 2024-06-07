package app

import (
	"log"
	"net/http"
	"time"

	"github.com/KrisjanisP/personal-dashboard/internal"
	"github.com/KrisjanisP/personal-dashboard/internal/repository"
	"github.com/KrisjanisP/personal-dashboard/web/templates/pages"
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
}

func NewApp(addr string) *App {
	app := &App{Addr: addr}

	sqliteDB := sqlx.MustConnect("sqlite3", "./data/sqlite3.db")

	app.sessionManager = scs.New()
	app.sessionManager.Store = sqlite3store.New(sqliteDB.DB)
	app.sessionManager.Lifetime = 24 * time.Hour

	app.userRepo = repository.NewUserRepository(sqliteDB)
	app.categoryRepo = repository.NewCategoryRepository(sqliteDB)

	return app
}

func (a *App) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("web/public"))).ServeHTTP)

	// routes that require auth cookie
	r.Group(func(r chi.Router) {
		r.Use(a.sessionManager.LoadAndSave)
		r.Use(a.Slow)
		r.Get("/login", a.viewAuthPage)

		r.Put("/login", a.login)
		r.Put("/logout", a.logout)

		r.Group(func(r chi.Router) {
			r.Use(a.AuthMiddleware)
			r.Put("/create/category", a.createCategory)
			r.Get("/", a.Home)
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

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int32)
	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		log.Println("error getting user by id:", err)
		if err := pages.ErrorPage("internal server error").Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	categories, err := a.categoryRepo.ListCategories(userID)
	if err != nil {
		log.Println("error getting categories:", err)
		if err := pages.ErrorPage("internal server error").Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := pages.HomePage(user, categories).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
