package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// Index handles the index page.
func Index(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "MyApp",
	}
	return c.Render(http.StatusOK, "index.html", data)
}
