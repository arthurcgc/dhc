package handlers

import (
	"net/http"

	"github.com/arthurcgc/dhc/internal/pkg/email"
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

// Email handles an email request from the client.
func Email(c echo.Context) error {
	contact := email.Contact{}
	if err := c.Bind(&contact); err != nil {
		return err
	}
	if err := email.Send(contact); err !=nil {
		return err
	}
	return c.String(http.StatusOK, "Email Sent")
}