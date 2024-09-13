// Package input implements input fields.
package input

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/form/validation/message"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/position"
	"github.com/jfbus/templ-components/components/size"
	"github.com/jfbus/templ-components/components/style"
)

const ()

type Type string

const (
	TypeEmail    Type = "email"
	TypeNumber   Type = "number"
	TypePassword Type = "password"
	TypeText     Type = "text"
	TypeURL      Type = "url"
)

// Defaults defines the default Color/Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"Class": {
			style.Color("bg-gray-50 border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"),
			style.Class("block w-full border rounded-lg"),
		},
		"ContainerClass": {},
		"IconClass": {
			style.Class("absolute inset-y-0 flex items-center pointer-events-none"),
		},
	},
	style.StyleDisabled: {
		"Class": {
			style.ReplaceColor("bg", "bg-gray-100 dark:bg-gray-700"),
			style.Add("cursor-not-allowed"),
		},
	},
	style.StyleValid: {
		"Class": {
			style.Color("bg-green-50 border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 focus:ring-green-500 focus:border-green-500 dark:bg-gray-700 dark:border-green-500"),
		},
		"IconClass": {
			style.Color("text-green-700 dark:text-green-500"),
		},
	},
	style.StyleInvalid: {
		"Class": {
			style.Add("bg-red-50 border-red-500 text-red-900 placeholder-red-700 focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 dark:text-red-500 dark:placeholder-red-500 dark:border-red-500"),
		},
		"IconClass": {
			style.Color("text-red-700 dark:text-red-500"),
		},
	},
}

// D is the definition for input fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Type is the input type (text, password, ...).
	Type Type
	// Style defines the style (style.StyleDefault, StyleValid or StyleInvalid).
	Style style.Style
	// Label is the input label.
	Label any
	// Value is the input value.
	Value string
	// Placeholder is the placeholder text displayed when no value is set.
	Placeholder string
	// Message adds a validation message below the field.
	// Just add &message.D{} to add automatic validation.
	//playground:import:github.com/jfbus/templ-components/components/form/validation/message
	//playground:default:&message.D{Message: "Validation message"}
	Message *message.D
	// Disabled disables the input.
	Disabled bool
	// Invalid marks the input as invalid.
	Invalid bool
	// Size defines the input size (size.S, size.Normal (default) or size.L).
	//playground:values:size.S,size.Normal,size.L
	Size size.Size
	// Loader displays a spinning loader when an HTMX action is triggered by the input.
	Loader bool
	// Icon displays an icon on the left side.
	//playground:values:icon.Flower
	Icon string
	// IconPosition can be position.Start (default) or position.End.
	//playground:values:position.Start,position.End
	IconPosition position.Position
	// Class overrides the default CSS class for the button.
	Class     style.D
	IconClass style.D
	// ContainerClass overrides the default CSS class for the div container.
	ContainerClass style.D
	// MessageClass overrides the default CSS class for the validation error container.
	MessageClass style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) style() style.Style {
	if def.Invalid {
		return def.Style | style.StyleInvalid
	}
	if def.Disabled {
		return def.Style | style.StyleDisabled
	}
	return def.Style
}

func (def D) label() label.D {
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID:      def.id(),
			InputName:    def.Name,
			Label:        l,
			Style:        def.style(),
			NoValidation: def.Message == nil,
		}
	case templ.Component:
		return label.D{
			InputID:      def.id(),
			InputName:    def.Name,
			Label:        l,
			Style:        def.style(),
			NoValidation: def.Message == nil,
		}
	case label.D:
		l.InputID = def.id()
		l.InputName = def.Name
		if l.Style == style.StyleDefault {
			l.Style = def.style()
		}
		l.NoValidation = def.Message == nil
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
	class := def.IconClass.CSSClass(Defaults, def.style(), "IconClass")
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
	class := def.Class.CSSClass(Defaults, def.style(), "Class")
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

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(Defaults, def.style(), "ContainerClass")
}

func (def D) inputClassInvalid() string {
	return def.Class.Delta(Defaults, def.Style, style.StyleInvalid, "Class")
}

func (def D) message() message.D {
	m := *def.Message
	m.InputName = def.Name
	m.Size = def.Size
	m.Style = def.style()
	return m
}
