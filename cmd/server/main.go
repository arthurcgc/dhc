package main

import (
	"log"

	"github.com/arthurcgc/dhc/internal/handlers"
	"github.com/arthurcgc/dhc/internal/renderer"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	// Create a new Echo instance
	e := echo.New()
	e.Renderer = renderer.New()
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register handler functions
	e.GET("/", handlers.Index)

	// Serve static files
	e.Static("/static", "static")

	// Start the server
	log.Fatal(e.Start(":8080"))
}
