// Package textarea implements textarea fields.
package textarea

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/helper"
	"github.com/jfbus/templ-components/input"
	"github.com/jfbus/templ-components/label"
	"github.com/jfbus/templ-components/position"
	"github.com/jfbus/templ-components/size"
)

// Defaults defines no defaults, you can set them, otherwise, input.Defaults are used.
var Defaults = D{}

// D is the definition for textarea fields.
type D struct {
	ID string
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Label any
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
	// Rows defines the number of rows
	Rows int
	// Disabled disables the input.
	Disabled bool
	// Size defines the input size (size.S, size.Normal (default) or size.L).
	Size size.Size
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Icon displays an icon on the left side.
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	IconPosition position.Position
	// Color overrides the default color CSS classes
	Color string
	// Class overrides the default CSS class for the textarea.
	Class string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) iconClass() string {
	class := "absolute inset-y-0 flex items-top pointer-events-none"
	switch def.Size {
	case size.S:
		class += " pt-2"
	case size.L:
		class += " pt-4"
	default:
		class += " pt-3.5"
	}
	switch {
	case def.Icon == "":
		return ""
	case def.IconPosition == position.End:
		return class + " end-0 pe-3.5"
	default:
		return class + " start-0 ps-3.5"
	}
}

func (def D) iconSize() size.Size {
	if def.Size == size.Inherit {
		return size.Normal
	}
	return def.Size
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
	switch {
	case def.Icon == "":
		return class
	case def.IconPosition == position.End:
		return class + " pe-10"
	default:
		return class + " ps-10"
	}
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
