package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SecureMiddleware returns an Echo middleware function for cors.
func SecureMiddleware(cfg config.SecurityConfig) echo.MiddlewareFunc {
	secureConfig := middleware.SecureConfig{
		XSSProtection:         cfg.XSSProtection,
		ContentTypeNosniff:    cfg.ContentTypeNosniff,
		XFrameOptions:         cfg.XFrameOptions,
		HSTSMaxAge:            cfg.HSTSMaxAge,
		ContentSecurityPolicy: cfg.ContentSecurityPolicy,
		ReferrerPolicy:        cfg.ReferrerPolicy,
	}

	return middleware.SecureWithConfig(secureConfig)
}
