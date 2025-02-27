package components

import "github.com/accentdesign/gtml"

func Submit(text string) *gtml.Element {
	return gtml.Button(gtml.Attrs{"class": "owl-button", "type": "submit"}, gtml.Text(text))
}
