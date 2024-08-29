// Package input implements input fields.
package input

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/helper"
	"github.com/jfbus/templ-components/label"
	"github.com/jfbus/templ-components/position"
	"github.com/jfbus/templ-components/size"
)

type Type string

const (
	TypeEmail    Type = "email"
	TypeNumber   Type = "number"
	TypePassword Type = "password"
	TypeText     Type = "text"
	TypeURL      Type = "url"
)

// Defaults defines the default Color/Class. They are overriden by D.Color/D.Class.
var Defaults = D{
	Color: "bg-gray-50 border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500",
	Class: "block w-full border rounded-lg",
}

// D is the definition for input fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Type Type
	// Label is the input label.
	Label any
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
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
	// Class overrides the default CSS class for the button.
	Class string
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
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

func (def D) id() string {
	if def.ID != "" {
		return def.ID
	}
	return def.Name
}

func (def D) inputType() string {
	if def.Type == "" {
		return string(TypeText)
	}
	return string(def.Type)
}

func (def D) iconClass() string {
	class := "absolute inset-y-0 flex items-center pointer-events-none"
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
	class := helper.IfEmpty(def.Class, Defaults.Class)
	class += " " + helper.IfEmpty(def.Color, Defaults.Color)
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
