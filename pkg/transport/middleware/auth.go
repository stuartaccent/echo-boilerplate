package middleware

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

// Authenticated middleware function to ensure the user is logged in; redirects to login if not.
func Authenticated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			setCurrentUser(c)

			if c.Get("user") == nil {
				cc := c.(*CustomContext)
				if cc.IsHTMXRequest() {
					cc.HTMXRedirect("/auth/login")
					return c.NoContent(http.StatusNoContent)
				} else {
					return c.Redirect(http.StatusFound, "/auth/login")
				}
			}

			return next(c)
		}
	}
}

// setCurrentUser sets the current active user in the context.
func setCurrentUser(c echo.Context) {
	cc := c.(*CustomContext)

	userIDInterface, ok := cc.Session.Values["user_id"]
	if !ok {
		return
	}

	userID, ok := userIDInterface.([16]byte)
	if !ok {
		return
	}

	user, err := cc.Queries.GetUserByID(c.Request().Context(), pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil || !user.IsActive {
		return
	}

	AddLogAttr(c, slog.String("user", user.Email))

	c.Set("user", user)
}
