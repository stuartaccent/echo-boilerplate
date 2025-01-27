package components

import (
	. "github.com/accentdesign/gtml"
)

func LoginForm(csrfToken, err string) *Element {
	return Form(
		Attrs{
			"class":     "w-full max-w-[350px] grid gap-10",
			"id":        "login-form",
			"method":    "post",
			"hx-post":   "/auth/login",
			"hx-select": "#login-form",
			"hx-swap":   "outerHTML",
		},
		Div(
			NA,
			H1(Attrs{"class": "owl-h2"}, Text("Login")),
			P(Attrs{"class": "text-muted-foreground"}, Text("Access your account.")),
		),
		Csrf(csrfToken),
		FormInput("email", "Email", "email", "", "", "", true),
		FormInput("password", "Password", "password", "", "", err, true),
		Submit("Login"),
	)
}
