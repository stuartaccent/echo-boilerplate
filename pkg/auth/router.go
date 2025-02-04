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

// LoginCredentials used in the login validation
type LoginCredentials struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

// Router create a new Router.
func Router(e *echo.Echo) {
	g := e.Group("/auth")
	{
		g.GET("/login", loginForm)
		g.POST("/login", login, middleware.AllowContentTypeForm)
		g.GET("/logout", logout)
	}
}

// loginForm get the login form
func loginForm(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cc.Session.Options.MaxAge = -1
	if err := cc.Session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return cc.RenderComponent(http.StatusOK, pages.Login(pages.LoginData{
		CsrfToken: c.Get("csrf").(string),
	}))
}

// login the user from the login form then redirect to home
func login(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	ctx := c.Request().Context()

	invalid := func() error {
		return cc.RenderComponent(http.StatusOK, pages.Login(pages.LoginData{
			Error:     "invalid email address or password",
			CsrfToken: c.Get("csrf").(string),
		}))
	}

	var credentials LoginCredentials
	if err := c.Bind(&credentials); err != nil {
		return invalid()
	}
	if err := c.Validate(credentials); err != nil {
		return invalid()
	}

	email := strings.ToLower(credentials.Email)
	user, err := cc.Queries.GetUserByEmail(ctx, email)
	if err != nil || !user.IsActive {
		return invalid()
	}

	password := []byte(credentials.Password)
	if !CheckPassword(user.HashedPassword, password) {
		return invalid()
	}

	cc.Session.Values["user_id"] = user.ID.Bytes
	if err = cc.Session.Save(c.Request(), c.Response()); err != nil {
		return invalid()
	}

	cc.HTMX.SetRedirect("/")

	return c.NoContent(http.StatusNoContent)
}

// logout the user then redirect to login
func logout(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	cc.Session.Options.MaxAge = -1
	if err := cc.Session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, "/auth/login")
}
