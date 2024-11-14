package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware returns an Echo middleware function for cors.
func CORSMiddleware(cfg config.SecurityConfig) echo.MiddlewareFunc {
	corsConfig := middleware.CORSConfig{
		AllowOrigins: cfg.AllowedHosts,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}

	return middleware.CORSWithConfig(corsConfig)
}
