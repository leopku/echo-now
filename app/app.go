package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

var app *echo.Echo

func Build() (e *echo.Echo) {
	e = echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if app == nil {
		app = Build()
	}
	app.ServeHTTP(w, r)
}
