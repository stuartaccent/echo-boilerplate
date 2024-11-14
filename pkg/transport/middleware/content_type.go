package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var AllowContentTypeForm = AllowContentType("application/x-www-form-urlencoded")

// AllowContentType enforces a whitelist of request Content-Types,
// otherwise responds with a 415 Unsupported Media Type status.
func AllowContentType(contentTypes ...string) echo.MiddlewareFunc {
	allowedContentTypes := make(map[string]struct{}, len(contentTypes))
	for _, c := range contentTypes {
		allowedContentTypes[strings.TrimSpace(strings.ToLower(c))] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := strings.ToLower(strings.TrimSpace(c.Request().Header.Get(echo.HeaderContentType)))
			if _, ok := allowedContentTypes[s]; ok {
				return next(c)
			}
			return c.NoContent(http.StatusUnsupportedMediaType)
		}
	}
}
