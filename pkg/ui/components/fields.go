package components

import (
	"github.com/a-h/templ"
	. "github.com/accentdesign/gtml"
)

func Csrf(value string) *Element {
	return Input(Attrs{"type": "hidden", "name": "_csrf", "value": value})
}

func FormInput(name, label, inputType, placeholder, helpText, errorText string, required bool) *Element {
	return Div(
		Attrs{"class": "owl-form-field"},
		Label(Attrs{"class": "owl-label", "for": name}, Text(label)),
		Input(Attrs{
			"class":       "owl-input",
			"id":          name,
			"name":        name,
			"placeholder": templ.KV(placeholder, placeholder != ""),
			"required":    required,
			"type":        inputType,
		}),
		If(errorText != "", P(Attrs{"class": "owl-form-field-error"}, Text(errorText))),
		If(helpText != "" && errorText == "", P(Attrs{"class": "owl-form-field-description"}, Text(helpText))),
	)
}
