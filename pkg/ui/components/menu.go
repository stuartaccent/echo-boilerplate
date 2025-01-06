package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func menuItem(label, url string) templ.Component {
	return gtml.A(gtml.Attrs{"class": "owl-dropdown-menu-item", "href": url, "role": "menuitem"}, gtml.Text(label))
}

func menuLabel(label string) templ.Component {
	return gtml.Div(gtml.Attrs{"class": "owl-dropdown-menu-label"}, gtml.Text(label))
}

func menuSeparator() templ.Component {
	return gtml.Div(gtml.Attrs{"class": "owl-dropdown-menu-separator", "role": "presentation"})
}

func UserMenu(opened bool) templ.Component {
	url := "/auth/user-menu"
	if !opened {
		url += "?open"
	}

	button := gtml.Button(
		gtml.Attrs{
			"class":      "owl-button owl-button-ghost",
			"hx-get":     url,
			"hx-trigger": templ.KV("click from:body, load delay:5s", opened),
		},
		gtml.Span(gtml.NA, gtml.Text("My Account")),
		Icon("chevron-down", ""),
	)

	menu := gtml.Div(
		gtml.Attrs{"class": "owl-dropdown-menu", "hx-target": "this", "hx-swap": "outerHTML"},
		button,
	)

	if opened {
		menu.AddChildren(
			gtml.Div(
				gtml.Attrs{"class": "owl-dropdown-menu-content owl-open right-0", "role": "menu"},
				menuLabel("My Account"),
				menuSeparator(),
				menuItem("Logout", "/auth/logout"),
			),
		)
	}

	return menu
}
