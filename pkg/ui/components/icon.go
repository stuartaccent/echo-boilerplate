package components

import (
	"github.com/a-h/templ"
	. "github.com/accentdesign/gtml"
)

func Icon(icon, class string) *Element {
	return &Element{
		Tag:   "owl-icon",
		Attrs: Attrs{"icon": icon, "class": templ.KV(class, class != "")},
	}
}
