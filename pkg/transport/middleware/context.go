package middleware

import (
	"echo.go.dev/pkg/storage/db/dbx"
	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	HTMX     *HTMX
	Postgres *pgxpool.Pool
	Queries  *dbx.Queries
	Session  *sessions.Session
}

func (c *CustomContext) RenderComponent(statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	return c.HTML(statusCode, buf.String())
}

// Context middleware func to define a custom context.
func Context(postgres *pgxpool.Pool) echo.MiddlewareFunc {
	queries := dbx.New(postgres)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			htmx := &HTMX{Request: c.Request(), Response: c.Response()}
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
				HTMX:     htmx,
				Postgres: postgres,
				Queries:  queries,
				Session:  sess,
			}
			return next(cc)
		}
	}
}
