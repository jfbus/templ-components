// Package radio implements radio fields.
package radio

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"radio": {
			style.StyleDefault: {
				style.Set("flex items-center"),
			},
			style.StyleDisabled: {
				style.Add("cursor-not-allowed"),
			},
		},
		"radio/input": {
			style.StyleDefault: {
				style.Set("w-4 h-4 focus:ring-2"),
			},
		},
	})
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
	defaults := def.LabelClass.WithDefault(def.style(), "radio/label")
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
	return def.ContainerClass.CSSClass(def.style(), "radio")
}

func (def D) inputClass() string {
	return def.InputClass.CSSClass(def.style(), "radio/input")
}
