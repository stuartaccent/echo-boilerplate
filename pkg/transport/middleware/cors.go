package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS returns an Echo middleware function for cors.
func CORS() echo.MiddlewareFunc {
	cfg := config.GetConfig()
	corsConfig := middleware.CORSConfig{
		AllowOrigins: cfg.Security.AllowedHosts,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}

	return middleware.CORSWithConfig(corsConfig)
}
