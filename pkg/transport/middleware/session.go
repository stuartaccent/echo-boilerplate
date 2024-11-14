package middleware

import (
	"echo.go.dev/pkg/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// SessionMiddleware returns an Echo middleware function for the session.
func SessionMiddleware(cfg config.SessionConfig) echo.MiddlewareFunc {
	sessionStore := sessions.NewCookieStore(cfg.KeyBytes(), cfg.EncKeyBytes())

	sessionStore.Options = &sessions.Options{
		Path:     cfg.Path,
		Domain:   cfg.Domain,
		MaxAge:   cfg.MaxAge,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
		SameSite: cfg.SameSite,
	}

	return session.Middleware(sessionStore)
}
