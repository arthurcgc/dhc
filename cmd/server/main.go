package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arthurcgc/dhc/internal/handlers"
	"github.com/arthurcgc/dhc/internal/renderer"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8888"

func main() {
	// Create a new Echo instance
	e := echo.New()
	e.Renderer = renderer.New()
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register handler functions
	e.GET("/", handlers.Index)
	e.POST("/email", handlers.Email)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(200, "up")
	})

	// Serve static files
	e.Static("/static", "static")

	// Set port
	serverPort := defaultPort
	if newPort := os.Getenv("PORT"); newPort != "" {
		serverPort = newPort
	}

	// Start the server
	log.Fatal(e.Start(fmt.Sprintf(":%s", serverPort)))
}
