package components

import (
	"github.com/a-h/templ"
	. "github.com/accentdesign/gtml"
)

func MenuItem(label, url string) *Element {
	return A(Attrs{"class": "owl-dropdown-menu-item", "href": url, "role": "menuitem"}, Text(label))
}

func MenuLabel(label string) *Element {
	return Div(Attrs{"class": "owl-dropdown-menu-label"}, Text(label))
}

func MenuSeparator() *Element {
	return Div(Attrs{"class": "owl-dropdown-menu-separator", "role": "presentation"})
}

func UserMenu(opened bool) *Element {
	url := "/auth/user-menu"
	if !opened {
		url += "?open"
	}

	return Div(
		Attrs{"class": "owl-dropdown-menu", "hx-target": "this", "hx-swap": "outerHTML"},
		Button(
			Attrs{
				"class":      "owl-button owl-button-ghost",
				"hx-get":     url,
				"hx-trigger": templ.KV("click from:body, load delay:5s", opened),
			},
			Span(NA, Text("My Account")),
			Icon("chevron-down", ""),
		),
		If(opened, Div(
			Attrs{"class": "owl-dropdown-menu-content owl-open right-0", "role": "menu"},
			MenuLabel("My Account"),
			MenuSeparator(),
			MenuItem("Logout", "/auth/logout"),
		)),
	)
}
