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

// Email handles an email request from the client.
func Email(c echo.Context) error {
	type Contact struct {
		Name string 	`json:"name"`
		Email string 	`json:"email"`
		Phone string 	`json:"phone"`
	}

	contact := Contact{}
	if err := c.Bind(&contact); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, contact)
}