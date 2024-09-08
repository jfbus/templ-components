// Package label implements labels.
package label

import (
	"github.com/jfbus/templ-components/components/style"
)

const StyleInline style.Style = 2

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
}

// D is the label definition.
type D struct {
	InputID string
	Style   style.Style
	Label   any
	Hide    bool
	Class   style.D
}

func (def D) class() string {
	return def.Class.CSSClass(Defaults, def.Style, "Class")
}
