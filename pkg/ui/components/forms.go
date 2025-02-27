package components

import "github.com/accentdesign/gtml"

func LoginForm(csrfToken, err string) *gtml.Element {
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
			gtml.P(gtml.Attrs{"class": "text-muted-foreground"}, gtml.Text("Access your account.")),
		),
		Csrf(csrfToken),
		FormInput("email", "Email", "email", "", "", "", true),
		FormInput("password", "Password", "password", "", "", err, true),
		Submit("Login"),
	)
}
