package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"library"
)

const PORT = 42069 // Nod to prime

type Pages map[string]*library.Page

var pages Pages = make(map[string]*library.Page)

func SetUpPages() Pages {
	if len(pages) == 0 {
		pages["index"] = library.NewPage("Index", "index", "World")
	}

	return pages
}

func main() {
	// Set up pages
	SetUpPages()

	// Create a new instance of Echo
	e := echo.New()
	e.Renderer = library.NewTemplates()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", indexHandler)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}

// handler
func indexHandler(c echo.Context) error {
	page := pages["index"]

	return c.Render(http.StatusOK, page.TemplateName, page)
}
