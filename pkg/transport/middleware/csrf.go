package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CSRF returns an Echo middleware function for csrf.
func CSRF() echo.MiddlewareFunc {
	cfg := config.GetConfig()
	csrfPaths := map[string]struct{}{
		"/auth/login": {},
	}

	csrfConfig := middleware.CSRFConfig{
		CookieName:     "_csrf",
		CookiePath:     cfg.Session.Path,
		CookieDomain:   cfg.Session.Domain,
		CookieSecure:   cfg.Session.Secure,
		CookieHTTPOnly: cfg.Session.HttpOnly,
		CookieSameSite: cfg.Session.SameSite,
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
