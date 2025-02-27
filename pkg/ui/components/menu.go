package components

import "github.com/accentdesign/gtml"

func MenuItem(label, url string) *gtml.Element {
	return gtml.A(gtml.Attrs{"class": "owl-dropdown-menu-item", "href": url, "role": "menuitem"}, gtml.Text(label))
}

func MenuLabel(label string) *gtml.Element {
	return gtml.Div(gtml.Attrs{"class": "owl-dropdown-menu-label"}, gtml.Text(label))
}

func MenuSeparator() *gtml.Element {
	return gtml.Div(gtml.Attrs{"class": "owl-dropdown-menu-separator", "role": "presentation"})
}

func UserMenu() *gtml.Element {
	return gtml.Div(gtml.Attrs{
		"class":  "owl-dropdown-menu",
		"x-data": "menu",
	},
		gtml.Button(gtml.Attrs{
			"class":  "owl-button owl-button-ghost",
			"x-ref":  "button",
			"@click": "toggle",
		},
			gtml.Span(gtml.NA, gtml.Text("My Account")),
			Icon("chevron-down", ""),
		),
		gtml.Div(gtml.Attrs{
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
