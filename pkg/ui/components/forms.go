package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func LoginForm(csrfToken, err string) templ.Component {
	return gtml.Form(
		gtml.Attrs{
			"class":     "w-full max-w-[350px] grid gap-10",
			"id":        "login-form",
			"method":    "post",
			"hx-post":   "/auth/login",
			"hx-select": "#login-form",
			"hx-swap":   "outerHTML",
		},
		gtml.Div(
			gtml.NA,
			gtml.H1(gtml.Attrs{"class": "owl-h2"}, gtml.Text("Login")),
			gtml.P(gtml.Attrs{"class": "text-gray-500"}, gtml.Text("Access your account.")),
		),
		Csrf(csrfToken),
		Input("email", "Email", "email", "", "", "", true),
		Input("password", "Password", "password", "", "", err, true),
		Submit("Login"),
	)
}
