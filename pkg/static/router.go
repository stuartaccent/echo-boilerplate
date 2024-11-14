package static

import (
	"embed"
	"github.com/labstack/echo/v4"
	"io/fs"
	"log"
)

var (
	//go:embed public/*
	Public embed.FS
)

// Router create a new static router.
func Router(e *echo.Echo) {
	e.StaticFS("/static", staticFS())
}

// staticFS returns the static file system.
func staticFS() fs.FS {
	s, err := fs.Sub(Public, "public")
	if err != nil {
		log.Fatalf("Unable to load static files: %v", err)
	}
	return s
}
