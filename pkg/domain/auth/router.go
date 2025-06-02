package auth

import (
	"echo.go.dev/pkg/transport/middleware"
	"echo.go.dev/pkg/ui/pages"
	"encoding/gob"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func init() {
	gob.Register([16]byte{})
}

// Router create a new Router.
func Router(e *echo.Echo) {
	g := e.Group("/auth")
	{
		g.GET("/login", loginHandler)
		g.POST("/login", loginPostHandler, middleware.AllowContentTypeForm)
		g.GET("/logout", logoutHandler)
	}
}

func loginHandler(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cc.Session.Options.MaxAge = -1
	if err := cc.Session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return cc.RenderComponent(http.StatusOK, pages.Login(pages.LoginProps{
		Csrf: c.Get("csrf").(string),
	}))
}

func loginPostHandler(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := c.Request().Context()

	invalid := func() error {
		return cc.RenderComponent(http.StatusOK, pages.Login(pages.LoginProps{
			Csrf:  c.Get("csrf").(string),
			Error: "invalid email or password",
		}))
	}

	var credentials struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
	}

	if err := c.Bind(&credentials); err != nil {
		return invalid()
	}
	if err := c.Validate(credentials); err != nil {
		return invalid()
	}

	email := strings.ToLower(credentials.Email)
	authUser, err := cc.Queries.GetUserByEmail(ctx, email)
	if err != nil || !authUser.IsActive {
		return invalid()
	}

	password := []byte(credentials.Password)
	if !CheckPassword(authUser.HashedPassword, password) {
		return invalid()
	}

	cc.Session.Values["user_id"] = authUser.ID.Bytes
	if err = cc.Session.Save(c.Request(), c.Response()); err != nil {
		return invalid()
	}

	cc.HTMXRedirect("/")

	return c.NoContent(http.StatusNoContent)
}

func logoutHandler(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cc.Session.Options.MaxAge = -1
	if err := cc.Session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/auth/login")
}
