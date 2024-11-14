package middleware

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// RateLimitMiddleware returns an Echo middleware function for rate limiting.
func RateLimitMiddleware() echo.MiddlewareFunc {
	ratePaths := map[string]struct{}{
		"POST:/auth/login": {},
	}

	rateStore := echomiddleware.NewRateLimiterMemoryStoreWithConfig(
		echomiddleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(10),
			Burst:     10,
			ExpiresIn: 1 * time.Minute,
		},
	)

	rateConfig := echomiddleware.RateLimiterConfig{
		Skipper: func(c echo.Context) bool {
			key := c.Request().Method + ":" + c.Path()
			if _, ok := ratePaths[key]; ok {
				return false
			}
			return true
		},
		Store: rateStore,
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			return ctx.RealIP(), nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, map[string]string{
				"error": "Rate limit exceeded. Please try again later.",
			})
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "Too many requests. Please slow down.",
			})
		},
	}

	return echomiddleware.RateLimiterWithConfig(rateConfig)
}
