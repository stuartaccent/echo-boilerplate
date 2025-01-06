package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Icon(icon, class string) templ.Component {
	svg := gtml.Element{
		Tag:   "owl-icon",
		Attrs: gtml.Attrs{"icon": icon, "class": templ.KV(class, class != "")},
	}
	return &svg
}
