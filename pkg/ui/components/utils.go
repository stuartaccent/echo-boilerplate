package components

import (
	. "github.com/accentdesign/gtml"
)

func If(condition bool, element *Element) *Element {
	if condition {
		return element
	}
	return Empty()
}
