// Package label implements labels.
package label

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

const StyleInline style.Style = 1 << 8

func init() {
	style.SetDefaults(style.Defaults{
		"label": {
			style.StyleDefault: {
				style.Set("block mb-2 text-sm font-medium"),
			},
			StyleInline: {
				style.Set("ms-2 text-sm font-medium"),
			},
		},
	})
}

// D is the label definition.
type D struct {
	// InputID is the ID of the corresponding input.
	InputID string
	// InputName is the Name of the corresponding input
	// (necessary for the validation code).
	InputName string
	// Style defines the label style.
	Style style.Style
	// Label can either be a string or a templ.Component.
	Label any
	// Hide hides the label.
	Hide bool
	// Class overrides the CSS classes.
	Class style.D
	// NoValidation removes the validation code.
	NoValidation bool
	// Attributes defines additional attributes.
	Attributes templ.Attributes
}

func (def D) class() string {
	return def.Class.CSSClass(def.Style, "label")
}

func (def D) classInvalid() string {
	return def.Class.Delta(def.Style, def.Style|style.StyleInvalid, "label")
}
