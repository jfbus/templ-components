// Package checkbox implements checkbox fields.
package checkbox

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/label"
	"github.com/jfbus/templui/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"checkbox": {
			style.Default: {
				style.Set("flex items-center"),
			},
		},
		"checkbox/input": {
			style.Default: {
				style.Set("ms-2 w-4 h-4 focus:ring-2 rounded"),
			},
			style.Disabled: {
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
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"checkbox":       style.D{style.Add("...")},
	// 		"checkbox/input": style.D{style.Add("...")},
	// 		"checkbox/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	Attributes  templ.Attributes
}

func (def D) style() style.Style {
	if def.Disabled {
		return style.Default | style.Disabled
	}
	return style.Default
}

func (def D) label() label.D {
	defaults := style.Custom{
		"label": style.Compute(def.style(), "checkbox/label", def.CustomStyle),
	}
	switch l := def.Label.(type) {
	case string:
		return label.D{
			InputID:     def.id(),
			Style:       label.StyleInline,
			Label:       l,
			CustomStyle: defaults,
		}
	case templ.Component:
		return label.D{
			InputID:     def.id(),
			Style:       label.StyleInline,
			Label:       l,
			CustomStyle: defaults,
		}
	case label.D:
		l.InputID = def.id()
		l.CustomStyle = l.CustomStyle.AddBefore(defaults)
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

func (def D) class(k string) string {
	return style.CSSClass(def.style(), k, def.CustomStyle)
}
