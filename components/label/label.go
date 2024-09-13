// Package label implements labels.
package label

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/style"
)

const StyleInline style.Style = 1 << 8

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Color("text-gray-900 dark:text-white"),
			style.Class("block mb-2 text-sm font-medium"),
		},
	},
	StyleInline: {
		"Class": {
			style.Class("ms-2 text-sm font-medium"),
		},
	},
	style.StyleValid: {
		"Class": {
			style.Color("text-green-700 dark:text-green-500"),
		},
	},
	style.StyleInvalid: {
		"Class": {
			style.Color("text-red-700 dark:text-red-500"),
		},
	},
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
	return def.Class.CSSClass(Defaults, def.Style, "Class")
}

func (def D) classInvalid() string {
	return def.Class.Delta(Defaults, def.Style, style.StyleInvalid, "Class")
}
