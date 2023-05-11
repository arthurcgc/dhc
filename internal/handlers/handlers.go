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

func Email(c echo.Context) error {
	type Contact struct {
		Name string 	`json:"name"`
		Email string 	`json:"email"`
		Message string 	`json:"message"`
	}

	contact := Contact{}
	if err := c.Bind(&contact); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Mensagem enviada com sucesso")
}