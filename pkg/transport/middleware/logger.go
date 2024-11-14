package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
	"os"
)

const (
	customAttributesCtxKey = "slog.custom_attributes"
)

func Logger() echo.MiddlewareFunc {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	loggerMiddleware := middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogHost:         true,
		LogURI:          true,
		LogRoutePath:    true,
		LogURIPath:      true,
		LogRemoteIP:     true,
		LogReferer:      true,
		LogLatency:      true,
		LogResponseSize: true,
		LogUserAgent:    true,
		LogStatus:       true,
		HandleError:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ctx := context.Background()
			level := slog.LevelInfo
			msg := "request"
			customAttrs, ok := c.Get(customAttributesCtxKey).([]slog.Attr)
			if !ok {
				customAttrs = make([]slog.Attr, 0)
			}
			attrs := []slog.Attr{
				slog.String("method", v.Method),
				slog.String("host", v.Host),
				slog.String("uri", v.URI),
				slog.String("route_path", v.RoutePath),
				slog.String("uri_path", v.URIPath),
				slog.String("remote_ip", v.RemoteIP),
				slog.String("referer", v.Referer),
				slog.Duration("latency", v.Latency),
				slog.Int64("response_size", v.ResponseSize),
				slog.String("user_agent", v.UserAgent),
				slog.Int("status", v.Status),
			}

			attrs = append(attrs, customAttrs...)

			if v.Error != nil {
				level = slog.LevelError
				msg = "request_error"
				attrs = append(attrs, slog.String("err", v.Error.Error()))
			}

			logger.LogAttrs(ctx, level, msg, attrs...)

			return nil
		},
	})

	return loggerMiddleware
}

func AddLogAttr(c echo.Context, attr slog.Attr) {
	attrs, ok := c.Get(customAttributesCtxKey).([]slog.Attr)
	if !ok {
		c.Set(customAttributesCtxKey, []slog.Attr{attr})
		return
	}
	c.Set(customAttributesCtxKey, append(attrs, attr))
}
