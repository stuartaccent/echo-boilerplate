package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Csrf(value string) *gtml.Element {
	return gtml.Input(gtml.Attrs{"type": "hidden", "name": "_csrf", "value": value})
}

func FormInput(name, label, inputType, placeholder, helpText, errorText string, required bool) *gtml.Element {
	return gtml.Div(
		gtml.Attrs{"class": "owl-form-field"},
		gtml.Label(gtml.Attrs{"class": "owl-label", "for": name}, gtml.Text(label)),
		gtml.Input(gtml.Attrs{
			"class":       "owl-input",
			"id":          name,
			"name":        name,
			"placeholder": templ.KV(placeholder, placeholder != ""),
			"required":    required,
			"type":        inputType,
		}),
		gtml.If(errorText != "", gtml.P(gtml.Attrs{"class": "owl-form-field-error"}, gtml.Text(errorText))),
		gtml.If(helpText != "" && errorText == "", gtml.P(gtml.Attrs{"class": "owl-form-field-description"}, gtml.Text(helpText))),
	)
}
