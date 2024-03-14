package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"library"
)

const PORT = 42069 // Nod to prime
var pages = library.GetPages()

func main() {
	// Create a new instance of Echo
	e := echo.New()
	e.Renderer = library.NewTemplates()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

// handler
func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index", pages["index"])
}
