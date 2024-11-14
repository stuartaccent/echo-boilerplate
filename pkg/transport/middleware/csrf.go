package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CSRFMiddleware returns an Echo middleware function for csrf.
func CSRFMiddleware(cfg config.SessionConfig) echo.MiddlewareFunc {
	csrfPaths := map[string]struct{}{
		"/auth/login": {},
	}

	csrfConfig := middleware.CSRFConfig{
		CookieName:     "_csrf",
		CookiePath:     cfg.Path,
		CookieDomain:   cfg.Domain,
		CookieSecure:   cfg.Secure,
		CookieHTTPOnly: cfg.HttpOnly,
		CookieSameSite: cfg.SameSite,
		TokenLookup:    "form:_csrf",
		ContextKey:     "csrf",
		Skipper: func(c echo.Context) bool {
			if _, ok := csrfPaths[c.Path()]; ok {
				return false
			}
			return true
		},
	}

	return middleware.CSRFWithConfig(csrfConfig)
}
