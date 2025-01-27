package components

import (
	. "github.com/accentdesign/gtml"
)

func PageHeader() *Element {
	return Header(
		NA,
		Div(
			Attrs{"class": "container mx-auto flex p-5 items-center"},
			A(Attrs{"class": "owl-h3 mr-auto", "href": "/"}, Text("Echo Boilerplate")),
			UserMenu(false),
		),
	)
}
