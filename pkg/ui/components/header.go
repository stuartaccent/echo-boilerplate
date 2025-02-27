package components

import "github.com/accentdesign/gtml"

func PageHeader() *gtml.Element {
	return gtml.Header(
		gtml.NA,
		gtml.Div(
			gtml.Attrs{"class": "container mx-auto flex p-5 items-center"},
			gtml.A(gtml.Attrs{"class": "owl-h3 mr-auto", "href": "/"}, gtml.Text("Echo Boilerplate")),
			UserMenu(),
		),
	)
}
