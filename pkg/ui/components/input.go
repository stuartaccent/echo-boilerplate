package components

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Csrf(value string) templ.Component {
	return gtml.Input(gtml.Attrs{"type": "hidden", "name": "_csrf", "value": value})
}

func Input(id, label, fieldType, placeholder, helpText, errorText string, required bool) templ.Component {
	cmp := gtml.Div(
		gtml.Attrs{"class": "owl-form-field"},
		gtml.Label(gtml.Attrs{"class": "owl-label", "for": id}, gtml.Text(label)),
		gtml.Input(gtml.Attrs{
			"class":       "owl-input",
			"id":          id,
			"name":        id,
			"placeholder": templ.KV(placeholder, placeholder != ""),
			"required":    required,
			"type":        fieldType,
		}),
	)
	if errorText != "" {
		cmp.AddChildren(gtml.P(gtml.Attrs{"class": "owl-form-field-error"}, gtml.Text(errorText)))
	} else if helpText != "" {
		cmp.AddChildren(gtml.P(gtml.Attrs{"class": "owl-form-field-description"}, gtml.Text(helpText)))
	}
	return cmp
}
