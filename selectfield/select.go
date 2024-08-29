// Package selectfield implements select fields.
// `select` being a reserved keyword, the package cannot be named select.
package selectfield

import (
	"github.com/jfbus/templ-components/helper"
	"github.com/jfbus/templ-components/input"
	"github.com/jfbus/templ-components/label"
	"github.com/jfbus/templ-components/selectfield/option"
	"github.com/jfbus/templ-components/size"
)

// Defaults defines the default CSS classes for select.
// If no values are defines, input.Defaults is used.
var Defaults = D{}

// D is the select definition.
type D struct {
	// ID is the tag id.
	ID string
	// Name is the tag name.
	Name string
	// Label is the label (either a string or a label.D).
	Label any
	// Options is the list of options.
	Options []option.D
	// Select is the selected value.
	Selected string
	// Size is the size.
	Size size.Size
	// Color overrides the default color CSS classes
	Color string
	// Class overrides the default CSS class for the select.
	Class string
}

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
}

func (def D) label() label.D {
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID: def.id(),
			Label:   l,
		}
	case label.D:
		l.InputID = def.id()
		return l
	default:
		return label.D{}
	}
}

func (def D) inputClass() string {
	class := helper.IfEmpty(def.Class, Defaults.Class, input.Defaults.Class)
	class += " " + helper.IfEmpty(def.Color, Defaults.Color, input.Defaults.Color)
	switch def.Size {
	case size.S:
		class += " p-2 text-xs"
	case size.L:
		class += " p-4 text-base"
	default:
		class += " p-2.5 text-sm"
	}
	return class
}
