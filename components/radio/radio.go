// Package radio implements radio fields.
package radio

import (
	"github.com/a-h/templ"
	"github.com/jfbus/templui/components/label"
	"github.com/jfbus/templui/components/style"
)

func init() {
	style.SetDefaults(style.Defaults{
		"radio": {
			style.Default: {
				style.Set("flex items-center"),
			},
			style.Disabled: {
				style.Add("cursor-not-allowed"),
			},
		},
		"radio/input": {
			style.Default: {
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
	// CustomStyle defines a custom style.
	// 	style.Custom{
	// 		"radio":       style.D{style.Add("...")},
	// 		"radio/input": style.D{style.Add("...")},
	// 		"radio/label": style.D{style.Add("...")},
	//	}
	CustomStyle style.Custom
	// Attributes stores additional attributes (e.g. HTMX attributes).
	Attributes templ.Attributes
}

func (def D) style() style.Style {
	if def.Disabled {
		return style.Default | style.Disabled
	}
	return style.Default
}

func (def D) label() label.D {
	cc := style.Custom{
		"label": style.Compute(def.style(), "radio/label", def.CustomStyle),
	}
	switch l := def.Label.(type) {
	case string:
		return label.D{
			Style:       label.StyleInline,
			InputID:     def.id(),
			Label:       l,
			CustomStyle: cc,
		}
	case templ.Component:
		return label.D{
			Style:       label.StyleInline,
			InputID:     def.id(),
			Label:       l,
			CustomStyle: cc,
		}
	case label.D:
		l.InputID = def.id()
		l.CustomStyle = l.CustomStyle.AddBefore(cc)
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
