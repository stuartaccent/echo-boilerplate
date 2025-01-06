package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Submit(text string) templ.Component {
	return gtml.Button(gtml.Attrs{"class": "owl-button", "type": "submit"}, gtml.Text(text))
}
