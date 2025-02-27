package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Icon(icon, class string) *gtml.Element {
	return gtml.New("owl-icon", gtml.Attrs{"icon": icon, "class": templ.KV(class, class != "")})
}
