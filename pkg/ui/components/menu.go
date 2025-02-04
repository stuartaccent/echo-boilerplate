package components

import (
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

func UserMenu() *Element {
	return Div(Attrs{
		"class":  "owl-dropdown-menu",
		"x-data": "menu",
	},
		Button(Attrs{
			"class":  "owl-button owl-button-ghost",
			"x-ref":  "button",
			"@click": "toggle",
		},
			Span(NA, Text("My Account")),
			Icon("chevron-down", ""),
		),
		Div(Attrs{
			"class":               "owl-dropdown-menu-content",
			"role":                "menu",
			"x-anchor.bottom-end": "$refs.button",
			"x-show":              "open",
			"x-transition":        true,
			"x-cloak":             true,
			"@click.outside":      "close",
		},
			MenuLabel("My Account"),
			MenuSeparator(),
			MenuItem("Logout", "/auth/logout"),
		),
	)
}
