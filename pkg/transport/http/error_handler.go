package http

import (
	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/ui/pages"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	c.Logger().Error(err)

	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	message := "An unexpected error occurred on the server."
	switch code {
	case http.StatusNotFound:
		message = "The requested URL was not found on this server."
	}

	cc := middleware.CustomContext{Context: c}
	if err := cc.RenderComponent(code, pages.Error(code, http.StatusText(code), message)); err != nil {
		c.Logger().Error(err)
	}
}
