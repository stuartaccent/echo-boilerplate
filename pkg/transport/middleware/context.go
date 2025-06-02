package middleware

import (
	"context"
	"echo.go.dev/pkg/storage/db/dbx"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Postgres *pgxpool.Pool
	Queries  *dbx.Queries
	Session  *sessions.Session
}

func (c *CustomContext) RenderComponent(statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	ctx := context.WithValue(c.Request().Context(), "Reverse", c.Reverse)
	if err := t.Render(ctx, buf); err != nil {
		return err
	}

	return c.HTML(statusCode, buf.String())
}

func (c *CustomContext) Reverse(url string, params ...interface{}) string {
	return c.Echo().Reverse(url, params...)
}

func (c *CustomContext) IsHTMXRequest() bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func (c *CustomContext) IsHTMXBoosted() bool {
	return c.Request().Header.Get("HX-Boosted") == "true"
}

func (c *CustomContext) HTMXRedirect(url string) {
	c.Response().Header().Set("HX-Redirect", url)
}

func (c *CustomContext) HTMXPushUrl(url string) {
	c.Response().Header().Set("HX-Push-Url", url)
}

func (c *CustomContext) HTMXRefresh() {
	c.Response().Header().Set("HX-Refresh", "true")
}

func (c *CustomContext) HTMXTrigger(content string) {
	c.Response().Header().Set("HX-Trigger", content)
}

func (c *CustomContext) HTMXTriggerAfterSettle(content string) {
	c.Response().Header().Set("HX-Trigger-After-Settle", content)
}

func (c *CustomContext) HTMXTriggerAfterSwap(content string) {
	c.Response().Header().Set("HX-Trigger-After-Swap", content)
}

// Context middleware func to define a custom context.
func Context(postgres *pgxpool.Pool) echo.MiddlewareFunc {
	queries := dbx.New(postgres)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			// if there is a problem with the session try to reset it
			if err != nil {
				sess.Options.MaxAge = -1
				if err := sess.Save(c.Request(), c.Response()); err != nil {
					return err
				}
			}
			cc := &CustomContext{
				Context:  c,
				Postgres: postgres,
				Queries:  queries,
				Session:  sess,
			}
			return next(cc)
		}
	}
}
