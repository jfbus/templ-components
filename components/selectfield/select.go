// Package selectfield implements select fields.
// `select` being a reserved keyword, the package cannot be named select.
package selectfield

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/input"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/selectfield/option"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

// Defaults defines the default CSS classes for select.
// If no values are defines, input.Defaults is used.
var Defaults = input.Defaults

// D is the select definition.
type D struct {
	// ID is the tag id.
	ID string
	// Name is the tag name.
	Name string
	// Label is the label (either a string or a label.D).
	Label any
	// Options is the list of options.
	//playground:import:github.com/jfbus/templ-components/components/selectfield/option
	//playground:default:[]option.D{{Label:"Select a value"},{Label:"Option 1"},{Label:"Option 2"}}
	Options []option.D
	// Select is the selected value.
	Selected string
	// Size is the size.
	Size size.Size
	// Class overrides the default CSS class for the select.
	Class      style.D
	Attributes templ.Attributes
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
	class := def.Class.CSSClass(style.Default(Defaults, style.StyleDefault, "Class"))
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
