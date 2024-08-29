// Package label implements labels.
package label

import (
	"github.com/jfbus/templ-components/components/helper"
)

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = D{
	Color: "text-gray-900 dark:text-white",
	Class: "block mb-2 text-sm font-medium",
}

// D is the label definition.
type D struct {
	InputID string
	Label   string
	Hide    bool
	Color   string
	Class   string
}

func (def D) class() string {
	return helper.IfEmpty(def.Class, Defaults.Class) + " " + helper.IfEmpty(def.Color, Defaults.Color)
}
