package components

import (
	. "github.com/accentdesign/gtml"
)

func Submit(text string) *Element {
	return Button(Attrs{"class": "owl-button", "type": "submit"}, Text(text))
}
