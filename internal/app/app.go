package app

import (
	"net/http"

	"github.com/labstack/echo"
)

type App struct {
	Addr string
}

func NewApp(addr string) *App {
	return &App{Addr: addr}
}

func (a *App) ListenAndServe() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(a.Addr))
}
