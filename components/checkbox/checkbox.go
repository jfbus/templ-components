// Package checkbox implements checkbox fields.
package checkbox

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templ-components/components/label"
	"github.com/jfbus/templ-components/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"checkbox": {
			style.StyleDefault: {
				style.Set("flex items-center"),
			},
		},
		"checkbox/input": {
			style.StyleDefault: {
				style.Set("ms-2 w-4 h-4 focus:ring-2 rounded"),
			},
			style.StyleDisabled: {
				style.Add("cursor-not-allowed"),
			},
		},
	})
}

// D is the definition for checkbox fields.
type D struct {
	// ID is the input id (Name if not set).
	ID string
	// Name is the input name.
	Name string
	// Label is the input label (either a string, a label.D or a templ.Component).
	//playground:edit:input
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
	defaults := def.LabelClass.WithDefault(def.style(), "checkbox/label")
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID: def.id(),
			Style:   label.StyleInline,
			Label:   l,
			Class:   defaults,
		}
	case templ.Component:
		return label.D{
			InputID: def.id(),
			Style:   label.StyleInline,
			Label:   l,
			Class:   defaults,
		}
	case label.D:
		l.InputID = def.id()
		if len(l.Class) == 0 {
			l.Class = defaults
		}
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
	return def.ContainerClass.CSSClass(def.style(), "checkbox")
}

func (def D) inputClass() string {
	return def.InputClass.CSSClass(def.style(), "checkbox/input")
}
