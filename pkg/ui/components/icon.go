package components

import (
	"github.com/a-h/templ"
	. "github.com/accentdesign/gtml"
)

func Icon(icon, class string) *Element {
	return New("owl-icon", Attrs{"icon": icon, "class": templ.KV(class, class != "")})
}
