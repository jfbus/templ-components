// Package label implements labels.
package label

import (
	"github.com/jfbus/templ-components/components/style"
)

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			Color: "text-gray-900 dark:text-white",
			Class: "block mb-2 text-sm font-medium",
		},
	},
}

// D is the label definition.
type D struct {
	InputID string
	Label   any
	Hide    bool
	Class   style.D
}

func (def D) class() string {
	return def.Class.CSSClass(style.Default(Defaults, style.StyleDefault, "Class"))
}
