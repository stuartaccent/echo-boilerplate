package domain

import (
	"echo.go.dev/pkg/domain/auth"
	"echo.go.dev/pkg/domain/home"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	auth.Router(e)
	home.Router(e)
}
