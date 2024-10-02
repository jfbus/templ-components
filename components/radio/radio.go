// Package radio implements radio fields.
package radio

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/style"
)

// Defaults defines the default Color/Class.
var Defaults = style.Defaults{
	style.StyleDefault: {
		"ContainerClass": style.D{
			style.Class("flex items-center"),
		},
		"InputClass": style.D{
			style.Class("w-4 h-4 focus:ring-2"),
			style.Color("text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:bg-gray-700 dark:border-gray-600"),
		},
	},
	style.StyleDisabled: {
		"LabelClass": style.D{
			style.ReplaceColor("text", "text-gray-400 dark:text-gray-500"),
		},
		"InputClass": style.D{
			style.Add("cursor-not-allowed"),
		},
	},
}

// D is the definition for radio fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Label is the input label (either a string, a label.D or a templ.Component).
	Label any
	// Value is the input value.
	Value   string
	Checked bool
	// Disabled disables the input.
	Disabled bool
	// ContainerClass overrides the class of the div container.
	ContainerClass style.D
	// InputClass overrides the class of the input tag.
	InputClass style.D
	// LabelClass overrides the class of the label tag.
	LabelClass style.D
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) style() style.Style {
	if def.Disabled {
		return style.StyleDefault | style.StyleDisabled
	}
	return style.StyleDefault
}

func (def D) label() label.D {
	defaults := def.LabelClass.WithDefault(Defaults, def.style(), "LabelClass")
	switch l := def.Label.(type) {
	case string:
		return label.D{
			Style:   label.StyleInline,
			InputID: def.id(),
			Label:   l,
			Class:   defaults,
		}
	case templ.Component:
		return label.D{
			Style:   label.StyleInline,
			InputID: def.id(),
			Label:   l,
			Class:   defaults,
		}
	case label.D:
		l.InputID = def.id()
		l.Class = append(defaults, l.Class...)
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

func (def D) containerClass() string {
	return def.ContainerClass.CSSClass(Defaults, def.style(), "ContainerClass")
}

func (def D) inputClass() string {
	return def.InputClass.CSSClass(Defaults, def.style(), "InputClass")
}
